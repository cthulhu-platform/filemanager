package service

// FileManagerService defines the interface for file manager operations
type FileManagerService interface {
	HandleDiagnoseMessage(transactionID, operation, message string) error
}
