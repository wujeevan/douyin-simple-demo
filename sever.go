package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/handler"
	"github.com/wujeevan/douyinv0/repository"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	if err := repository.Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/douyin/feed", handler.QueryFeedVideo)
	r.GET("/upload/:filename", func(ctx *gin.Context) {
		filepath := "." + ctx.Request.URL.String()
		filename := ctx.Query("filename")
		if strings.Contains(filename, ".mp4") {
			ctx.Header("content-type", "video/mp4")
		} else if strings.Contains(filename, ".jpg") {
			ctx.Header("content-type", "image/jpeg")
		}
		ctx.File(filepath)
	})
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
