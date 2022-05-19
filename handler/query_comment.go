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
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		SendFailResponse(ctx, err)
		return
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		SendFailResponse(ctx, err)
		return
	}
	if err := service.ProcessComment(videoId, actionType, token, content); err != nil {
		SendFailResponse(ctx, err)
	} else {
		SendSuccessResponse(ctx)
	}
}

func QueryCommentList(ctx *gin.Context) {
	token := ctx.Query("token")
	videoIdStr := ctx.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		ctx.JSON(200, &VideoCommentResponse{
			Code:        -1,
			Msg:         err.Error(),
			CommentList: []NULL{},
		})
	} else {
		commentList, err := service.QueryCommentList(token, videoId)
		if err != nil {
			ctx.JSON(200, &VideoCommentResponse{
				Code:        -1,
				Msg:         err.Error(),
				CommentList: []NULL{},
			})
		}
		ctx.JSON(200, &VideoCommentResponse{
			Code:        0,
			Msg:         "success",
			CommentList: commentList,
		})
	}
}
