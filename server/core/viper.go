package core

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func Viper() (v *viper.Viper) {
	v = viper.New()

	src, err := os.Getwd()
	if err != nil {
		log.Fatal("获取当前文件路径失败")
		return
	}

	v.AddConfigPath(fmt.Sprintf("%s/conf/", src))

	v.SetConfigName("app")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("获取配置文件失败, err:%v", err)
	}

	log.Print("--获取配置文件成功--")

	return
}
