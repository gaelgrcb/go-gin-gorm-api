package main

import (
	"CRUD/config"
	"CRUD/pkg/database"
	"CRUD/pkg/response"
	"log"
	"net/http"

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

	port := ":" + cfg.AppPort

	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Error initializing the server: %v", err)
	}
}
