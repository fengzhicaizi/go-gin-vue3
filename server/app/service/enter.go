package service

import (
	"github.com/fengzhicaizi/gin-vue3/app/service/common"
	"github.com/fengzhicaizi/gin-vue3/app/service/system"
)

type ServiceGroup struct {
	System system.ServiceGroup
	Common common.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
