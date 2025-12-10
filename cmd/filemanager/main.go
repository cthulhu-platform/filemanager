package main

import (
	"log"
	"path/filepath"

	"github.com/cthulhu-platform/common/pkg/env"
	"github.com/cthulhu-platform/filemanager/internal/pkg"
	"github.com/cthulhu-platform/filemanager/internal/repository"
	"github.com/cthulhu-platform/filemanager/internal/server"
	"github.com/cthulhu-platform/filemanager/internal/service"
)

func main() {
	// Load environment variables from .env file if it exists
	// Look for .env in the filemanager directory
	envPath := filepath.Join(".", ".env")
	env.SetupEnvFile(envPath)

	// Initialize repository
	r, err := repository.NewLocalRepository(pkg.STORAGE_PATH)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer r.Close()

	// Initialize service
	s := service.NewFileManagerService(r)

	// Configure RabbitMQ server
	cfg := &server.RMQServerConfig{
		User:           pkg.AMQP_USER,
		Password:       pkg.AMQP_PASS,
		Host:           pkg.AMQP_HOST,
		Port:           pkg.AMQP_PORT,
		VHost:          pkg.AMQP_VHOST,
		ConnectionName: "filemanager",
	}

	// Start RabbitMQ server
	server.ListenRMQ(s, cfg)
}
