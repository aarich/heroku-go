package api

import (
	"fmt"
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
	"github.com/aarich/heroku-go/pkg/util"
	"github.com/gin-gonic/gin"
)

func Generate(c *gin.Context) {
	a := app.GinApp{Context: c}

	file, image, err := c.Request.FormFile("image")

	if err != nil {
		log.Println(err)
		a.RespondError(http.StatusInternalServerError, errors.UNKNOWN, "")
		return
	}

	if image == nil {
		a.RespondError(http.StatusBadRequest, errors.INVALID_PARAMS, "Missing image.")
		return
	}
	imageName := getImageName(image.Filename)
	fullPath := getImageFullPath()
	savePath := getImagePath()
	src := fullPath + imageName

	if !checkImageExt(imageName) || !checkImageSize(file) {
		a.RespondError(http.StatusBadRequest, errors.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, "")
		return
	}

	err = checkImage(fullPath)
	if err != nil {
		log.Println(err)
		a.RespondError(http.StatusInternalServerError, errors.ERROR_UPLOAD_CHECK_IMAGE_FAIL, "")
		return
	}

	err = a.Context.SaveUploadedFile(image, src)
	if err != nil {
		log.Println(err)
		a.RespondError(http.StatusInternalServerError, errors.ERROR_UPLOAD_SAVE_IMAGE_FAIL, "")
		return
	}

	a.RespondSuccess(http.StatusOK, map[string]string{
		"image_url":      getImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}

func getImageName(name string) string {
	ext := filepath.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
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
