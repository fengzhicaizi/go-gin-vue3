package v1

import (
	"github.com/fengzhicaizi/gin-vue3/app/api/v1/system"
)

type ApiGroup struct {
	System system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
