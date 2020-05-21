package common

import (
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// 日志管理
func Logger(logPath string) *logrus.Logger {
	logPath = filepath.FromSlash(logPath)

	path, filename := filepath.Split(logPath)
	logger := logger(env("log/"+path), filename)

	return logger
}

// 日志定义与分割
func logger(path, filename string) *logrus.Logger {
	l := logrus.New()

	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	Error(err)

	l.Out = src
	l.SetLevel(logrus.DebugLevel)

	if _, err = os.Stat(path); err != nil {
		Error(os.MkdirAll(path, os.ModePerm))
	}

	logPath := path + "/" + filename
	logWriter, err := rotatelogs.New(
		logPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	Error(err)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	l.AddHook(lfHook)

	return l
}
