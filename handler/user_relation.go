package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
)

type FollowResponse struct {
	Code     int         `json:"status_code"`
	Msg      string      `json:"status_msg"`
	UserList interface{} `json:"user_list"`
}

func ProcessFollowUser(ctx *gin.Context) {
	token := ctx.Query("token")
	followedIdStr := ctx.Query("to_user_id")
	actionTypeStr := ctx.Query("action_type")
	followedId, err := strconv.ParseInt(followedIdStr, 10, 64)
	if err != nil {
		SendFailResponse(ctx, err)
		return
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		SendFailResponse(ctx, err)
		return
	}
	if err := service.ProcessFollowUser(token, actionType, followedId); err != nil {
		SendFailResponse(ctx, err)
	} else {
		SendSuccessResponse(ctx)
	}
}

func QueryFollowList(ctx *gin.Context) {
	token := ctx.Query("token")
	userList, err := service.QueryFollowList(token)
	if err != nil {
		ctx.JSON(200, &FollowResponse{
			Code:     -1,
			Msg:      err.Error(),
			UserList: []NULL{},
		})
	}
	ctx.JSON(200, &FollowResponse{
		Code:     0,
		Msg:      "success",
		UserList: userList,
	})
}

func QueryFollowerList(ctx *gin.Context) {
	token := ctx.Query("token")
	userList, err := service.QueryFollowerList(token)
	if err != nil {
		ctx.JSON(200, &FollowResponse{
			Code:     -1,
			Msg:      err.Error(),
			UserList: []NULL{},
		})
	}
	ctx.JSON(200, &FollowResponse{
		Code:     0,
		Msg:      "success",
		UserList: userList,
	})
}
