package facade

import (
	"guildchat-oss/pkg/newer"
)

type RedisLog struct {
	logger *newer.RedisLogger
}

func NewRedisLog(path string) *RedisLog {

	return &RedisLog{
		logger: newer.NewRedisLogger(path),
	}
}

func (rlog RedisLog) Info(msg string, info map[string]string) {

	rlog.logger.Info(msg, info)
}

func (rlog RedisLog) Error(msg string, info map[string]string) {
	rlog.logger.Error(msg, info)
}
