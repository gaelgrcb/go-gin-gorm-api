package main

import (
	"log"
	"net/http"

	"github.com/gaelgrcb/go-gin-gorm-api/config"
	"github.com/gaelgrcb/go-gin-gorm-api/pkg/database"
	"github.com/gaelgrcb/go-gin-gorm-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	_, err = database.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("database error: %v", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		response.Ok(c, "Successful Response", nil)
	})

	router.GET("/api", func(c *gin.Context) {
		response.Ok(c, "Successful Response", nil)
	})

	port := ":" + cfg.AppPort

	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Error initializing the server: %v", err)
	}
}
