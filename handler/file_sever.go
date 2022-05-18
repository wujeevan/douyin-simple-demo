package handler

import (
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func FileSever(ctx *gin.Context) {
	filepath := path.Join(".", ctx.Request.URL.String())
	filename := ctx.Query("filename")
	if strings.Contains(filename, ".mp4") {
		ctx.Header("content-type", "video/mp4")
	} else if strings.Contains(filename, ".jpg") {
		ctx.Header("content-type", "image/jpeg")
	}
	ctx.File(filepath)
}
