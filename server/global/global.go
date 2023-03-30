package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GVA_VP    *viper.Viper
	GVA_ZAP   *zap.Logger
	GVA_DB    *gorm.DB
	GVA_REDIS redis.Conn
)
