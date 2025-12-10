package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cthulhu-platform/common/pkg/messages"
	"github.com/cthulhu-platform/filemanager/internal/handlers"
	"github.com/cthulhu-platform/filemanager/internal/service"
	"github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
)

type RMQServerConfig struct {
	User           string
	Password       string
	Host           string
	Port           string
	VHost          string
	ConnectionName string
}

// ListenRMQ sets up and starts a RabbitMQ consumer for the filemanager service
func ListenRMQ(s service.FileManagerService, cfg *RMQServerConfig) {
	// Create connection string
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.VHost,
	)

	// Create RabbitMQ connection with labeled connection name
	conn, err := rabbitmq.NewConn(
		connectionString,
		rabbitmq.WithConnectionOptionsLogging,
		rabbitmq.WithConnectionOptionsConfig(rabbitmq.Config{
			Properties: amqp091.Table{
				"connection_name": cfg.ConnectionName,
			},
		}),
	)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create consumer for diagnose messages
	consumer, err := rabbitmq.NewConsumer(
		conn,
		"filemanager.diagnose",
		rabbitmq.WithConsumerOptionsExchangeName(messages.DiagnoseExchange),
		rabbitmq.WithConsumerOptionsExchangeKind("topic"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsQueueDurable,
		rabbitmq.WithConsumerOptionsBinding(rabbitmq.Binding{
			RoutingKey: messages.TopicDiagnoseServicesAll,
		}),
	)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	// Start consuming messages
	go func() {
		if err := consumer.Run(handlers.HandleDiagnoseMessage(s)); err != nil {
			log.Fatalf("Consumer error: %v", err)
		}
	}()

	log.Printf("FileManager service started and listening for messages on exchange: %s", messages.DiagnoseExchange)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down gracefully...")
}
