package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Run starts the URL shortener server
func Run(port, host, baseURL, dbPath string) error {
	// Initialize database connection
	db, err := NewDB(dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	// Initialize server
	router := gin.Default()

	// Define routes
	router.POST("/shorten", shortenHandler(db))
	router.GET("/:key", redirectHandler(db))

	// Start server
	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Starting server on %s", addr)
	return router.Run(addr)
}
