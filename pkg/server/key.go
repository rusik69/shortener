package server

import "github.com/google/uuid"

// generateKey generates a random key for the URL
func generateKey() string {
	// Generate a random UUID and take first 8 characters
	// This gives us 16^8 = 4.3 billion possible combinations
	// while keeping URLs reasonably short
	id := uuid.New().String()
	return id[:8]
}
