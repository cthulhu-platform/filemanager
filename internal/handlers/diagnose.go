package handlers

import (
	"encoding/json"
	"log"

	"github.com/cthulhu-platform/common/pkg/messages"
	"github.com/cthulhu-platform/filemanager/internal/service"
	"github.com/wagslane/go-rabbitmq"
)

// HandleDiagnoseMessage processes diagnose messages from RabbitMQ
func HandleDiagnoseMessage(s service.FileManagerService) func(d rabbitmq.Delivery) rabbitmq.Action {
	return func(d rabbitmq.Delivery) rabbitmq.Action {
		// Unmarshal the message
		var msg messages.DiagnoseMessage
		if err := json.Unmarshal(d.Body, &msg); err != nil {
			log.Printf("Failed to unmarshal diagnose message: %v", err)
			return rabbitmq.NackRequeue
		}

		// Process the message through the service
		if err := s.HandleDiagnoseMessage(msg.TransactionID, msg.Operation, msg.Message); err != nil {
			log.Printf("Failed to handle diagnose message: %v", err)
			return rabbitmq.NackRequeue
		}

		log.Printf("Successfully processed diagnose message: %s", msg.TransactionID)
		return rabbitmq.Ack
	}
}

