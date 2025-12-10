package service

import (
	"log"

	"github.com/cthulhu-platform/filemanager/internal/repository"
)

type fileManagerService struct {
	repo repository.Repository
}

func NewFileManagerService(repo repository.Repository) FileManagerService {
	return &fileManagerService{
		repo: repo,
	}
}

func (s *fileManagerService) HandleDiagnoseMessage(transactionID, operation, message string) error {
	log.Printf("FileManager received diagnose message - TransactionID: %s, Operation: %s, Message: %s",
		transactionID, operation, message)

	// Handle the diagnose message
	// For now, just log it. Add more logic as needed.

	return nil
}
