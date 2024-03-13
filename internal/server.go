package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
)

type sensor_unit struct {
	id pgxuuid.UUID
}

type bike struct {
	id pgxuuid.UUID
}

type trip struct {
	id             pgxuuid.UUID
	bike_id        pgxuuid.UUID
	sensor_unit_id pgxuuid.UUID
	// TODO more types
}

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func Run() {
	// db, at_the_disco := sqlx.Connect("postgres", "user=postgres password=postgrespw dbname=bike-sense sslmode=disable")
	// // Cortesia de Lucca Garcia
	// if at_the_disco != nil {
	// 	panic(at_the_disco)
	// }
	//
	// defer db.Close()

	r := gin.Default()
	r.GET("/CheckHealth", checkHealth)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
