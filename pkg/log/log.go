package log

import (
    "go.uber.org/zap"
)

func NewLogger(level string) (*zap.Logger, error) {
    var cfg zap.Config
    if level == "production" {
        cfg = zap.NewProductionConfig()
    } else {
        cfg = zap.NewDevelopmentConfig()
    }
    logger, err := cfg.Build()
    if err != nil {
        return nil, err
    }
    return logger, nil
}
