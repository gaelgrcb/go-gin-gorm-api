package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
