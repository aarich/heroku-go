package app

import (
	"github.com/aarich/heroku-go/pkg/errors"
	"github.com/gin-gonic/gin"
)

type GinApp struct {
	Context *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (app *GinApp) Respond(httpCode, errCode int, data interface{}) {
	app.Context.IndentedJSON(httpCode, Response{errCode, errors.GetMessage(errCode), data})
}
