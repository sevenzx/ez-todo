package initialize

import (
	"github.com/sevenzx/eztodo/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

func Lumberjack() *lumberjack.Logger {
	// 可定制的输出目录。
	var logFilePath string
	dir := "./"
	logFilePath = dir + "/logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// 将文件名设置为日期
	logFileName := time.Now().Format(time.DateOnly) + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}
	// 提供压缩和删除
	c := config.Config.Lumberjack
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    c.MaxSize,    // 一个文件最大可达 20M。
		MaxBackups: c.MaxBackups, // 最多同时保存 5 个文件。
		MaxAge:     c.MaxAge,     // 一个文件最多可以保存 10 天。
		Compress:   c.Compress,   // 用 gzip 压缩。
	}
}
