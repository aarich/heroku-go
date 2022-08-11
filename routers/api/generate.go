package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aarich/heroku-go/pkg/app"
	"github.com/aarich/heroku-go/pkg/errors"
	"github.com/aarich/heroku-go/pkg/file"
	"github.com/aarich/heroku-go/pkg/settings"
	"github.com/aarich/heroku-go/pkg/stereogram"
	"github.com/aarich/heroku-go/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/harrydb/go/img/grayscale"
)

func Generate(c *gin.Context) {
	a := app.GinApp{Context: c}

	file, info, err := c.Request.FormFile("image")

	if err != nil {
		log.Println(err)
		a.RespondError(http.StatusInternalServerError, errors.UNKNOWN, "")
		return
	}

	if info == nil {
		a.RespondError(http.StatusBadRequest, errors.INVALID_PARAMS, "Missing image.")
		return
	}

	if !okContentType(info.Header.Get("Content-Type")) {
		a.RespondError(http.StatusBadRequest, errors.INVALID_PARAMS, "Image is wrong content type")
		return
	}

	bs, err := ioutil.ReadAll(file)

	if err != nil {
		a.RespondError(http.StatusBadRequest, errors.UNKNOWN, "error reading file")
		return
	}

	image, _, err := image.Decode(bytes.NewReader(bs))

	if err != nil {
		a.RespondError(http.StatusBadRequest, errors.UNKNOWN, "error decoding file")
		return
	}

	z := imageTo2dArray(&image)
	result := stereogram.Generate(z)

	s := &bytes.Buffer{}
	err = png.Encode(s, result)
	if err != nil {
		log.Fatalln(err)
		return
	}
	str := base64.StdEncoding.EncodeToString(s.Bytes())

	a.RespondSuccess(http.StatusOK, map[string]string{
		"url": "data:image/png;base64," + str,
	})
	//
	// imageName := getImageName(image.Filename)
	// fullPath := getImageFullPath()
	// savePath := getImagePath()
	// src := fullPath + imageName
	//
	// if !checkImageExt(imageName) || !checkImageSize(file) {
	// 	a.RespondError(http.StatusBadRequest, errors.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, "")
	// 	return
	// }
	//
	// err = checkImage(fullPath)
	// if err != nil {
	// 	log.Println(err)
	// 	a.RespondError(http.StatusInternalServerError, errors.ERROR_UPLOAD_CHECK_IMAGE_FAIL, "")
	// 	return
	// }
	//
	// err = a.Context.SaveUploadedFile(image, src)
	// if err != nil {
	// 	log.Println(err)
	// 	a.RespondError(http.StatusInternalServerError, errors.ERROR_UPLOAD_SAVE_IMAGE_FAIL, "")
	// 	return
	// }
	//
	// a.RespondSuccess(http.StatusOK, map[string]string{
	// 	"image_url":      getImageFullUrl(imageName),
	// 	"image_save_url": savePath + imageName,
	// })
}

func imageTo2dArray(im *image.Image) [][]float32 {
	gray := grayscale.Convert(*im, grayscale.ToGrayLuma)
	var ret [][]float32

	// For each pixel scale down the luminosity from 0-255 to 0-1
	for y := 0; y < gray.Bounds().Dy(); y++ {
		var row []float32
		for x := 0; x < gray.Bounds().Dx(); x++ {
			val := gray.GrayAt(x, y).Y
			row = append(row, float32(val)/255)
		}
		ret = append(ret, row)
	}

	return ret
}

func getImageName(name string) string {
	ext := filepath.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func okContentType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/gif"
}

func getImagePath() string {
	return settings.App.ImageSavePath
}

func getImageFullPath() string {
	return settings.App.RuntimeRootPath + getImagePath()
}

func checkImageExt(fileName string) bool {
	ext := strings.ToUpper(file.GetExt(fileName))
	for _, allowExt := range settings.App.ImageAllowedExts {
		if strings.ToUpper(allowExt) == ext {
			return true
		}
	}

	return false
}

func checkImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}

	return size <= settings.App.ImageMaxSize
}

func checkImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}

func getImageFullUrl(name string) string {
	return settings.App.PrefixUrl + "/" + getImagePath() + name
}
