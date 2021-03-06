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
	QueryUserVideo("fEWqWdVNwL62834d13d71qE0gdV5")
}
