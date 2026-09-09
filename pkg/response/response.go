package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status string `json:"status"` // "OK", "WARN", "ERR"
	Msg    string `json:"msg"`
	Data   any    `json:"data,omitempty"`
}

// Response: 200
func Ok(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, Response{
		Status: "OK",
		Msg:    message,
		Data:   data,
	})
}

func Warn(c *gin.Context, message string, detail any) {
	c.JSON(http.StatusBadRequest, Response{
		Status: "WARN",
		Msg:    message,
		Data:   detail,
	})
}

func Error(c *gin.Context, message string, err any) {
	c.JSON(http.StatusInternalServerError, Response{
		Status: "ERR",
		Msg:    message,
		Data:   err,
	})
}
