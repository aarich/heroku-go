package routers

import (
	"net/http"

	"github.com/aarich/heroku-go/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/generate", api.Generate)
	router.GET("/sample", api.Sample)

	return router
}
