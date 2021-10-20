package loggers

import "guildchat-oss/pkg/loggers/facade"

/*
* 通用info日志
 */
func LogInfo(path string, msg string, info map[string]string) {
	log := facade.NewZaplog(path)
	log.Info(msg, info)
}

/*
* 通用error日志
 */
func LogError(path string, msg string, info map[string]string) {
	log := facade.NewZaplog(path)
	log.Error(msg, info)
}
