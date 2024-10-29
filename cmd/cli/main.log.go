package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d",  "TipsDev😂", 24);

	// logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsDev"), zap.Int("age", 24));
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsDev😂"), zap.Int("age", 24));

	//Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production")

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// format log a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// timestamp -> date time format
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encodeConfig.TimeKey = "time"
	// from info -> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//"caller":"cli/main.log.go:24"
	encodeConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}