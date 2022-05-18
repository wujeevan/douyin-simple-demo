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
	followedId, err1 := strconv.ParseInt(followedIdStr, 10, 64)
	actionType, err2 := strconv.ParseInt(actionTypeStr, 10, 64)
	if err1 != nil || err2 != nil {
		ctx.JSON(200, &FollowResponse{
			Code:     -1,
			Msg:      err1.Error(),
			UserList: []NULL{},
		})
	} else {
		if err := service.ProcessFollowUser(token, actionType, followedId); err != nil {
			ctx.JSON(200, &FollowResponse{
				Code:     -1,
				Msg:      err1.Error(),
				UserList: []NULL{},
			})
		} else {
			ctx.JSON(200, &FollowResponse{
				Code: 0,
				Msg:  "success",
			})
		}
	}
}

func QueryFollowList(ctx *gin.Context) {
	token := ctx.Query("token")
	userList, err := service.QueryFollowList(token)
	if err != nil {
		ctx.JSON(200, &FollowResponse{
			Code: -1,
			Msg:  err.Error(),
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
			Code: -1,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(200, &FollowResponse{
		Code:     0,
		Msg:      "success",
		UserList: userList,
	})
}
