package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/controller"
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
	r.GET("/douyin/feed", func(ctx *gin.Context) {
		latest_time := ctx.Query("latest_time")
		token := ctx.Query("token")
		ctx.JSON(200, controller.QueryFeedVideo(latest_time, token))
	})
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
