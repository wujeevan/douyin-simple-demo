package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
)

type UserInfoResponse struct {
	Code int               `json:"status_code"`
	Msg  string            `json:"status_msg"`
	User *service.UserInfo `json:"user"`
}

func SendFailResponse(ctx *gin.Context, err error) {
	ctx.JSON(200, gin.H{
		"status_code": -1,
		"status_msg":  err.Error(),
	})
}

func SendSuccessResponse(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "success",
	})
}

func QueryUserInfo(ctx *gin.Context) {
	token := ctx.Query("token")
	userInfo, err := service.QueryUserByToken(token)
	if err != nil {
		SendFailResponse(ctx, err)
	} else {
		ctx.JSON(200, &UserInfoResponse{
			Code: 0,
			Msg:  "success",
			User: userInfo,
		})
	}
}
