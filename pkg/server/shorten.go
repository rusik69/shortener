package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// shortenHandler handles the shorten request
func shortenHandler(db *DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			URL string `json:"url" binding:"required,url"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Generate a unique key
		key := generateKey()

		// Store the URL in the database
		if err := db.Set(key, request.URL); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
			return
		}

		// Return the shortened URL
		c.JSON(http.StatusOK, gin.H{"short_url": fmt.Sprintf("%s/%s", c.Request.Host, key)})
	}
}
