package handler

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
	"github.com/wujeevan/douyinv0/utils"
)

func UploadVideo(ctx *gin.Context) {
	token := ctx.PostForm("token")
	file, err := ctx.FormFile("data")
	if err != nil {
		SendFailResponse(ctx, err)
	}
	filename := utils.GenerateFilename("mp4")
	filepath := path.Join("./upload", filename)
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		SendFailResponse(ctx, err)
	}
	if err := service.UploadVideo(token, filename); err != nil {
		SendFailResponse(ctx, err)
	} else {
		SendSuccessResponse(ctx)
	}
}
