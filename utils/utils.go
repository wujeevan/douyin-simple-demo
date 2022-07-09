package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var TokenValidTime = 7 * 24 * time.Hour

func RandStr(n int) string {
	b := make([]rune, n)
	rand.Seed(time.Now().Unix())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateFilename(suffix string) string {
	curYearDay := time.Now().YearDay()
	randStr := RandStr(8)
	filename := fmt.Sprintf("%s%x%s", randStr, curYearDay, suffix)
	return filename
}

func GenerateToken() string {
	curTime := time.Now().Unix()
	randStr := RandStr(20)
	token := fmt.Sprintf("%s%x%s", randStr[:10], curTime, randStr[10:])
	return token
}

func CheckSqlInjection(str string) error {
	if strings.ContainsAny(str, " &|^='<>/*%#") {
		return errors.New("invalid string")
	}
	return nil
}

func AddHostName(host, path string) string {
	return "http://" + host + path
}

func GenerateVideoCover(filepath string) (string, error) {
	path := strings.Split(filepath, ".")
	path[0] += "_cover"
	path[1] = "jpg"
	cover_filepath := strings.Join(path, ".")
	cmd := exec.Command("ffmpeg", "-y", "-ss", "1", "-i", "."+filepath, "-vframes", "1", "-vcodec", "mjpeg", "."+cover_filepath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return "", errors.New("cover of video generate failed")
	}
	return cover_filepath, nil
}
