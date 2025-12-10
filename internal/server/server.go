package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cthulhu-platform/common/pkg/messages"
	"github.com/cthulhu-platform/filemanager/internal/handlers"
	"github.com/cthulhu-platform/filemanager/internal/service"
	"github.com/wagslane/go-rabbitmq"
)

type RMQServer struct {
	Conn    *rabbitmq.Conn
	Service service.FileManagerService
}

// NewRMQServer creates a new RabbitMQ server instance
func NewRMQServer(conn *rabbitmq.Conn, s service.FileManagerService) *RMQServer {
	return &RMQServer{
		Conn:    conn,
		Service: s,
	}
}

// Start sets up and starts the diagnose consumer
func (s *RMQServer) Start() {
	// Create diagnose consumer - simplified to match library example pattern
	consumer, err := rabbitmq.NewConsumer(
		s.Conn,
		"filemanager.diagnose",
		rabbitmq.WithConsumerOptionsRoutingKey(messages.TopicDiagnoseServicesAll),
		rabbitmq.WithConsumerOptionsExchangeName(messages.DiagnoseExchange),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeKind("topic"),
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsQueueDurable,
		rabbitmq.WithConsumerOptionsConsumerName("filemanager_diagnose"),
	)
	if err != nil {
		log.Fatalf("Failed to create diagnose consumer: %v", err)
	}
	defer consumer.Close()

	log.Printf("FileManager service started and listening for messages on exchange: %s with routing key: %s",
		messages.DiagnoseExchange, messages.TopicDiagnoseServicesAll)

	// Setup graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Printf("Received signal: %v, stopping consumer...", sig)
		consumer.Close()
	}()

	// Block main thread - wait for messages (Run is blocking)
	log.Println("Starting consumer and waiting for messages...")
	if err := consumer.Run(handlers.HandleDiagnoseMessage(s.Service)); err != nil {
		log.Fatalf("Consumer error: %v", err)
	}

	log.Println("Shutting down gracefully...")
}
