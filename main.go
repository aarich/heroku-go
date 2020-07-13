package main

import (
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/aarich/heroku-go/pkg/settings"
	"github.com/aarich/heroku-go/pkg/util"
	"github.com/aarich/heroku-go/routers"
)

func init() {
	settings.Setup()
}

func main() {

	router := routers.InitRouter()

	port := util.GetEnv("PORT")
	router.Run(":" + port)
}
