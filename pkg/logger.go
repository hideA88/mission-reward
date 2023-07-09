package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func NewLogger(verbose bool) *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = localTimeEncoder
	if !verbose {
		encoderConfig.CallerKey = ""
		encoderConfig.LevelKey = ""
	}

	//	cfg := zap.NewProductionConfig()
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig = encoderConfig

	rowLogger, _ := cfg.Build()

	//nolint:errcheck
	defer rowLogger.Sync() // flushes buffer, if any

	return rowLogger.Sugar()
}

func localTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	local, _ := time.LoadLocation("Local")
	enc.AppendString(t.In(local).Format(time.RFC3339))
}
