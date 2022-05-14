package controller

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/wujeevan/douyinv0/repository"
)

func TestMain(m *testing.M) {
	if err := repository.Init(); err != nil {
		os.Exit(1)
	}
	m.Run()
}

func TestQueryFeedVideo(t *testing.T) {
	QueryFeedVideo(strconv.FormatInt(time.Now().UnixMilli(), 10), "hello")
}
