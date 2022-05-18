package repository

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := Init(); err != nil {
		os.Exit(1)
	}
	m.Run()
}

func TestVideoDao_QueryFeedVideoNoToken(t *testing.T) {
	QueryFeedVideo(1652313600, 1)
}

func TestQueryVideoById(t *testing.T) {
	QueryUserVideoById(1)
}

func TestQueryFavoriteVideo(t *testing.T) {
	QueryFavoriteVideo(1)
}
