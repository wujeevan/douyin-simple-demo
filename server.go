package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/go-simple-community/controller"
	"github.com/wujeevan/go-simple-community/repository"
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
	if err := repository.Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:topicId", func(c *gin.Context) {
		topicId := c.Param("topicId")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	r.POST("/community/page/post/post", func(c *gin.Context) {
		topicId := c.PostForm("topicId")
		content := c.PostForm("content")
		data := controller.InsertPost(topicId, content)
		c.JSON(200, data)
	})
	r.POST("/community/page/post/topic", func(c *gin.Context) {
		topicId := c.PostForm("title")
		content := c.PostForm("content")
		data := controller.InsertTopic(topicId, content)
		c.JSON(200, data)
	})
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
