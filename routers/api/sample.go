package api

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"net/http"

	"github.com/aarich/heroku-go/pkg/app"
	"github.com/aarich/heroku-go/pkg/errors"
	"github.com/aarich/heroku-go/pkg/images"
	"github.com/aarich/heroku-go/pkg/stereogram"
	"github.com/gin-gonic/gin"
)

func Sample(c *gin.Context) {
	size := image.Rect(0, 0, 300, 300)
	z := images.MakeDistanceMap(size, images.Square)
	result := stereogram.Generate(z)

	a := app.GinApp{c}

	s := &bytes.Buffer{}
	png.Encode(s, result)
	str := base64.StdEncoding.EncodeToString(s.Bytes())

	a.Respond(http.StatusOK, errors.SUCCESS, map[string]string{
		"result": str,
		"url":    "data:image/png;base64," + str,
	})
}
