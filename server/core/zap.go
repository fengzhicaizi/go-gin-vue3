package core

import (
	"log"

	"go.uber.org/zap"
)

func Zap() (z *zap.Logger) {
	z = zap.NewExample()
	log.Print("--初始化日志成功--")
	return
}
