package logger

import "go.uber.org/zap"

var logger *zap.SugaredLogger

// Should be called only from main!
func Init() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger = zapLogger.Sugar()
}

func Get() *zap.SugaredLogger {
	return logger
}

func Close() {
	logger.Sync()
}

