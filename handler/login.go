package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
)

type SignInResponse struct {
	Code   int    `json:"status_code"`
	Msg    string `json:"status_msg"`
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type SignUpResponse = SignInResponse

func SignIn(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	signInfo, err := service.SignIn(username, password)
	if err != nil {
		SendFailResponse(ctx, err)
	} else {
		ctx.JSON(200, &SignInResponse{
			Code:   0,
			Msg:    "sign in success",
			UserId: signInfo.UserId,
			Token:  signInfo.Token,
		})
	}
}

func SignUp(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	signInfo, err := service.SignUp(username, password)
	if err != nil {
		SendFailResponse(ctx, err)
	} else {
		ctx.JSON(200, &SignInResponse{
			Code:   0,
			Msg:    "sign up success",
			UserId: signInfo.UserId,
			Token:  signInfo.Token,
		})
	}
}
