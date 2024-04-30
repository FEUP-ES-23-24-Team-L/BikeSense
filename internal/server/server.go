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

	v1 := router.Group("/api/v1")
	{
		v1.GET("/check_health", routes.CheckHealth)
		v1.POST("/sensor_unit/register", routes.PostSensorUnit)
		v1.POST("/bike/register", routes.PostBike)
		v1.POST("/trip/register", routes.PostTrip)
		v1.POST("/trip/upload_data", routes.PostTripData)
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
