package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/cthulhu-platform/common/pkg/env"
	"github.com/cthulhu-platform/filemanager/internal/pkg"
	"github.com/cthulhu-platform/filemanager/internal/repository"
	"github.com/cthulhu-platform/filemanager/internal/server"
	"github.com/cthulhu-platform/filemanager/internal/service"
	"github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
)

func main() {
	// Load environment variables from .env file if it exists
	// Look for .env in the filemanager directory
	envPath := filepath.Join(".", ".env")
	if err := env.Init(envPath); err != nil {
		// Non-fatal: will fall back to environment variables
		log.Printf("Warning: Could not load .env file: %v\n", err)
	}

	// Initialize repository
	r, err := repository.NewLocalRepository(pkg.STORAGE_PATH)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer r.Close()

	// Initialize service
	fileService := service.NewFileManagerService(r)

	// Create RabbitMQ connection
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s%s",
		pkg.AMQP_USER,
		pkg.AMQP_PASS,
		pkg.AMQP_HOST,
		pkg.AMQP_PORT,
		pkg.AMQP_VHOST,
	)
	fmt.Println("Connection string: ", connectionString)

	conn, err := rabbitmq.NewConn(
		connectionString,
		rabbitmq.WithConnectionOptionsLogging,
		rabbitmq.WithConnectionOptionsConfig(rabbitmq.Config{
			Properties: amqp091.Table{
				"connection_name": "filemanager",
			},
		}),
	)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create and start RabbitMQ server
	s := server.NewRMQServer(conn, fileService)
	s.Start()
}
