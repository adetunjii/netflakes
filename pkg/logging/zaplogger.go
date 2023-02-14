package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SugarLogger interface {
	Infof(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
}

var _ SugarLogger = (*zap.SugaredLogger)(nil)

func NewZapSugarLogger() *zap.SugaredLogger {
	enconderConfig := newEncoderConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(enconderConfig)

	// create a new zap core that prints log out to the console
	logger := zap.New(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.WarnLevel),
	)
	return logger.Sugar()
}

func newEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	return config
}
