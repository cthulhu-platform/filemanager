package repository

import (
	"fmt"
	"os"
	"path/filepath"
)

type LocalRepository struct {
	storagePath string
}

func NewLocalRepository(storagePath string) (*LocalRepository, error) {
	// Create storage directory if it doesn't exist
	if err := os.MkdirAll(storagePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &LocalRepository{
		storagePath: storagePath,
	}, nil
}

func (r *LocalRepository) Close() error {
	// Nothing to close for local file system
	return nil
}

// GetStoragePath returns the storage path
func (r *LocalRepository) GetStoragePath() string {
	return r.storagePath
}

// EnsureDirectory ensures a directory exists within storage
func (r *LocalRepository) EnsureDirectory(dir string) error {
	fullPath := filepath.Join(r.storagePath, dir)
	return os.MkdirAll(fullPath, 0755)
}
