package core

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fengzhicaizi/gin-vue3/initialize"

	. "github.com/fengzhicaizi/gin-vue3/global"
)

func RunServer() {

	var (
		addr         = fmt.Sprintf(":%v", GVA_VP.GetString("server.HttpPort"))
		port         = GVA_VP.GetString("server.HttpPort")
		readTimeout  = time.Duration(GVA_VP.GetInt64("server.ReadTimeout")) * time.Second
		writeTimeout = time.Duration(GVA_VP.GetInt64("server.WriteTimeout")) * time.Second
		router       = initialize.Router()
	)

	s := http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("系统启动成功 监听端口: %v", port)
	log.Printf("swaggo地址: http://127.0.0.1:8001/swagger/index.html")

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("系统启动失败")
	}
}
