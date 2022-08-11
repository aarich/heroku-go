package routers

import (
	"github.com/aarich/heroku-go/routers/api"
	"github.com/aarich/heroku-go/routers/pages"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", pages.Index)

	router.GET("/api/generate", api.Generate)
	router.GET("/api/sample", api.Sample)

	return router
}
