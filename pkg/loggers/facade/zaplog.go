package facade

import (
	"go.uber.org/zap"
	"guildchat-oss/pkg/loggers/newer"
)

type ZapLog struct {
	logger *zap.Logger
}

func NewZaplog(path string) *ZapLog {
	return &ZapLog{
		logger: newer.NewZapLogger(path),
	}
}

func (zlog ZapLog) Info(msg string, info map[string]string) {
	zapSlice := make([]zap.Field, len(info))
	var fieldNum int
	for k, v := range info {
		zapSlice[fieldNum] = zap.String(k, v)
		fieldNum++
	}
	zlog.logger.Info(msg, zapSlice...)
}

func (zlog ZapLog) Error(msg string, info map[string]string) {
	zapSlice := make([]zap.Field, len(info))
	var fieldNum int
	for k, v := range info {
		zapSlice[fieldNum] = zap.String(k, v)
		fieldNum++
	}
	zlog.logger.Error(msg, zapSlice...)
}
