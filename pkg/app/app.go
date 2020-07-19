package app

import (
	"github.com/aarich/heroku-go/pkg/errors"
	"github.com/gin-gonic/gin"
)

type GinApp struct {
	Context *gin.Context
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (app *GinApp) RespondSuccess(httpCode int, data interface{}) {
	app.Context.IndentedJSON(httpCode, data)
}

func (app *GinApp) RespondError(httpCode, errCode int, extendedMessage string) {
	message := errors.GetMessage(errCode)
	if message != "" {
		message += " " + extendedMessage
	}

	app.Context.IndentedJSON(httpCode, ErrorResponse{errCode, message})
}
