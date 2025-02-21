package services

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, ErrorResponse{Error: err.Error()})
}
