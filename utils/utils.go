package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

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
	filename := fmt.Sprintf("%s%x.%s", randStr, curYearDay, suffix)
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
