
package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RespondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, ErrorResponse{Message: message})
}
