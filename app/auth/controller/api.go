package api

import (
	"database/sql"

	// "git.eways.dev/eways/service/storage"

	"github.com/gin-gonic/gin"
)

func Populate(r *gin.Engine, db *sql.DB) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	// api := r.Group("/api/stocks")
	// api.Use(SentryErrorReporter())

	// migrationRoutes(api, db)
}
