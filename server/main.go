package main

import (
	"github.com/fengzhicaizi/gin-vue3/core"
	"github.com/fengzhicaizi/gin-vue3/global"
	"github.com/fengzhicaizi/gin-vue3/initialize"
)

// 初始化
func init() {
	global.GVA_VP = core.Viper()      // 启动配置
	global.GVA_ZAP = core.Zap()       // 启动日志
	global.GVA_DB = initialize.Grom() // 启动数据库链接
	// global.GVA_REDIS = initialize.Redis() // 初始化redis
}

// @title 后台管理系统
// @version 1.0
// @description 后管平台
func main() {
	core.RunServer() // 启动服务器
}
