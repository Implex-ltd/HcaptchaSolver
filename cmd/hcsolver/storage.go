package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func CreateLogFile() *os.File {
	logFileName := fmt.Sprintf("../../assets/logs/%s", time.Now().Format("2006-01-02_15-04-05")+".json")

	err := os.MkdirAll("../../assets/logs/", os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

func LoadSettings() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder

	fileEncoder := zap.NewProductionEncoderConfig()
	fileEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),
			zapcore.AddSync(colorable.NewColorableStdout()),
			zapcore.DebugLevel,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(fileEncoder),
			zapcore.AddSync(CreateLogFile()),
			zapcore.DebugLevel,
		),
	)

	logger = zap.New(core)
}
