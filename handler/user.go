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

func QueryUserInfo(ctx *gin.Context) {
	token := ctx.Query("token")
	userInfo, err := service.QueryUserByToken(token)
	if err != nil {
		ctx.JSON(200, &UserInfoResponse{
			Code: -1,
			Msg:  err.Error(),
		})
	} else {
		ctx.JSON(200, &UserInfoResponse{
			Code: 0,
			Msg:  "success",
			User: userInfo,
		})
	}
}
