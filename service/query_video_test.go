package service

import (
	"os"
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
	QueryFeedVideo(time.Now().Unix(), "")
}

func TestQueryUserVideo(t *testing.T) {
	QueryUserVideo("30fe3c5cc4159001bc7966d619d2ec204f43a1bc695428ee180c83f9a2c")
}
