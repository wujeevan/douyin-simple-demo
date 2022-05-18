package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
)

type VideoCommentResponse struct {
	Code        int         `json:"status_code"`
	Msg         string      `json:"status_msg"`
	CommentList interface{} `json:"comment_list"`
}

func ProcessComment(ctx *gin.Context) {
	token := ctx.Query("token")
	videoIdStr := ctx.Query("video_id")
	actionTypeStr := ctx.Query("action_type")
	content := ctx.Query("comment_text")
	videoId, err1 := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, err2 := strconv.ParseInt(actionTypeStr, 10, 64)
	if err1 != nil || err2 != nil {
		ctx.JSON(200, &VideoCommentResponse{
			Code:        -1,
			Msg:         err1.Error(),
			CommentList: []NULL{},
		})
	} else {
		if err := service.ProcessComment(videoId, actionType, token, content); err != nil {
			ctx.JSON(200, &VideoCommentResponse{
				Code: -1,
				Msg:  err1.Error(),
			})
		} else {
			ctx.JSON(200, &VideoCommentResponse{
				Code: 0,
				Msg:  "success",
			})
		}
	}
}

func QueryCommentList(ctx *gin.Context) {
	token := ctx.Query("token")
	videoIdStr := ctx.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		ctx.JSON(200, &VideoCommentResponse{
			Code: -1,
			Msg:  err.Error(),
		})
	} else {
		commentList, err := service.QueryCommentList(token, videoId)
		if err != nil {
			ctx.JSON(200, &VideoCommentResponse{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		ctx.JSON(200, &VideoCommentResponse{
			Code:        0,
			Msg:         "success",
			CommentList: commentList,
		})
	}
}
