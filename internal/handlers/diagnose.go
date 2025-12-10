package handlers

import (
	"encoding/json"
	"log"

	"github.com/cthulhu-platform/common/pkg/messages"
	"github.com/cthulhu-platform/filemanager/internal/service"
	"github.com/wagslane/go-rabbitmq"
)

// HandleDiagnoseMessage processes diagnose messages from RabbitMQ
func HandleDiagnoseMessage(s service.FileManagerService) rabbitmq.Handler {
	return func(d rabbitmq.Delivery) rabbitmq.Action {
		log.Printf("Handler received message - RoutingKey: %s, Body: %s", d.RoutingKey, string(d.Body))
		
		// Unmarshal the message
		var msg messages.DiagnoseMessage
		if err := json.Unmarshal(d.Body, &msg); err != nil {
			log.Printf("Failed to unmarshal diagnose message: %v", err)
			return rabbitmq.NackRequeue
		}

		// Print the diagnosis check message
		log.Printf("Diagnosis check from routing key: %s", d.RoutingKey)
		log.Printf("Received message - TransactionID: %s, Operation: %s, Message: %s",
			msg.TransactionID, msg.Operation, msg.Message)

		// Process the message through the service
		if err := s.HandleDiagnoseMessage(msg.TransactionID, msg.Operation, msg.Message); err != nil {
			log.Printf("Failed to handle diagnose message: %v", err)
			return rabbitmq.NackRequeue
		}

		return rabbitmq.Ack
	}
}
