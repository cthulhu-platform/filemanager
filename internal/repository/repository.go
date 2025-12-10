package repository

// Repository defines the interface for file storage operations
type Repository interface {
	Close() error
	// Add more methods as needed for file operations
}
