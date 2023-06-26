package utils

import "github.com/gin-gonic/gin"

// NewError creates an error in the gin context.
//
// TODO: log error in this helper function
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	ctx.JSON(status, er)
}

// HTTPError is the error response for the API.
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
