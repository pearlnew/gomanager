package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var Log = zap.NewNop().Sugar()

func InitLogger(path string){
	var logger *zap.SugaredLogger
	if path == ""{
		logger = zap.NewExample().Sugar()
	} else {
		core := zapcore.NewCore(getEncoder(), getLogWriter(path), zapcore.InfoLevel)
		sugarLogger := zap.New(core, zap.AddCaller())
		logger = sugarLogger.Sugar()
	}
	defer logger.Sync()
	Log = logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	file, _ := os.Create(path)
	return zapcore.AddSync(file)
}