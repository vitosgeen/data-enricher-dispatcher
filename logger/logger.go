package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Fatal(args ...interface{})
	Println(args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
}

func NewLogger() Logger {
	logger := logrus.New()

	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Level = logrus.InfoLevel

	// save logs to file
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		logger.Fatal("Failed to open log file:", err)
	}
	logger.Out = file
	return logger
}
