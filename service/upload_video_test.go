package service

import (
	"testing"
)

func TestUploadVideo(t *testing.T) {
	UploadVideo("userpass", "./upload/test4.mp4")
}

func TestGenerateVideoCover(t *testing.T) {
	GenerateVideoCover("/upload/test4.mp4")
}
