package initialize

import (
	"log"

	. "github.com/fengzhicaizi/gin-vue3/global"

	"github.com/gomodule/redigo/redis"
)

func Redis() (r redis.Conn) {
	var err error
	r, err = redis.Dial("tcp", GVA_VP.GetString("redis.Host"), redis.DialPassword(GVA_VP.GetString("redis.PaasWord")))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("--redis 准备就绪,地址:%s --", GVA_VP.GetString("redis.Host"))
	return
}
