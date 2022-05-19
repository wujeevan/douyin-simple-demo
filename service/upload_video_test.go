package service

import (
	"testing"
)

func TestUploadVideo(t *testing.T) {
	UploadVideo("userpass", "./upload/test4.mp4")
}
