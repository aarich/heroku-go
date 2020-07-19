package api

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"log"
	"net/http"

	"github.com/aarich/heroku-go/pkg/app"
	"github.com/aarich/heroku-go/pkg/errors"
	"github.com/aarich/heroku-go/pkg/images"
	"github.com/aarich/heroku-go/pkg/stereogram"
	"github.com/gin-gonic/gin"
)

func Sample(c *gin.Context) {
	a := app.GinApp{Context: c}

	queryParams := c.Request.URL.Query()
	chosenType := queryParams.Get("type")
	if chosenType == "" {
		a.RespondError(http.StatusBadRequest, errors.INVALID_PARAMS, "Missing 'type' (options: 'square', 'steps').")
		return
	}

	size := image.Rect(0, 0, 300, 300)
	z := images.MakeDistanceMap(size, chosenType)
	result := stereogram.Generate(z)

	s := &bytes.Buffer{}
	err := png.Encode(s, result)
	if err != nil {
		log.Fatalln(err)
	}
	str := base64.StdEncoding.EncodeToString(s.Bytes())

	a.RespondSuccess(http.StatusOK, map[string]string{
		"url":  "data:image/png;base64," + str,
		"type": chosenType,
	})
}
