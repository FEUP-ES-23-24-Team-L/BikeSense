package server

import (
	"bikesense-web/internal/server/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func dbMiddleware(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("db", db) // Set the db connection in the context
		c.Next()
	}
}

func Run(db *gorm.DB) {
	if db == nil {
		panic("db is nil")
	}

	router := gin.Default()

	router.Use(dbMiddleware(db))

	router.GET("/CheckHealth", routes.CheckHealth)
	router.POST("/Trip", routes.PostTrip)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
