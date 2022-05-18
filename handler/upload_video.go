package handler

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
	"github.com/wujeevan/douyinv0/utils"
)

func UploadVideo(ctx *gin.Context) {
	token := ctx.PostForm("token")
	file, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
	}
	filename := utils.GenerateFilename("mp4")
	filepath := path.Join("./upload", filename)
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
	}
	if err := service.UploadVideo(token, filename); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"statud_code": 0,
			"status_msg":  "success",
		})
	}
}
