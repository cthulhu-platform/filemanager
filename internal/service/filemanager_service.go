package service

import (
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
	// Handle the diagnose message
	// The logging is done in the handler layer
	return nil
}
