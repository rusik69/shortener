package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// redirectHandler handles the redirect request
func redirectHandler(db *DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Param("key")
		url, err := db.Get(key)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.Redirect(http.StatusMovedPermanently, url)
	}
}
