package router

import (
	"github.com/fengzhicaizi/gin-vue3/app/router/example"
	"github.com/fengzhicaizi/gin-vue3/app/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
