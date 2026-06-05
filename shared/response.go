package shared

import (
	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, data any, error *AppError) {
	// TODO add check on error layer to not respond with any private error details
	if Conf.ISLocal == "true" {
		if error != nil {
			c.JSON(
				400,
				Response{
					Success: false,
					Message: "something went wrong",
					// Error:   &error.Error,
				},
			)
			return
		}
	} else {
		if error != nil {
			// err := errors.New(error.Message)
			c.JSON(
				400,
				Response{
					Success: false,
					Message: "something went wrong",
					// Error:   &err,
				},
			)
			return
		}
	}
	c.JSON(
		200,
		Response{
			Success: true,
			Message: "Success",
			Data:    data,
		},
	)

}

// func Fail[T any](c *gin.Context, error error) {
// 	c.JSON(
// 		400,
// 		Response[T]{
// 			Success: false,
// 			Message: "something went wrong",
// 			Error:   error,
// 		},
// 	)
// }
