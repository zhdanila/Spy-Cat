package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sca/internal/config"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	var level zapcore.Level
	err := level.UnmarshalText([]byte(cfg.Logger.LogLevel))
	if err != nil {
		return nil, err
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      false,
		Encoding:         cfg.Logger.Encoding,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      cfg.Logger.OutputPaths,
		ErrorOutputPaths: cfg.Logger.ErrorOutputPaths,
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
