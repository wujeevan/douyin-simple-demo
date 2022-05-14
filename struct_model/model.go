package main

import (
	"github.com/qmhball/db2gorm/gen"
)

func main() {
	dsn := "root:801925@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

	// tblName := "user_comment"
	// gen.GenerateOne(gen.GenConf{
	// 	Dsn:       dsn,
	// 	WritePath: "./model",
	// 	Stdout:    false,
	// 	Overwrite: true,
	// }, tblName)
	gen.GenerateAll(gen.GenConf{
		Dsn:       dsn,
		WritePath: "../model",
		Stdout:    false,
		Overwrite: true,
	})
}
