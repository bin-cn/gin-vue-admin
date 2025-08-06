package systask

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/router"
	"github.com/gin-gonic/gin"
)

type SysTaskPlugin struct {
}

func CreateSysTaskPlug() *SysTaskPlugin {

	return &SysTaskPlugin{}
}

func (*SysTaskPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitSysTaskRouter(group)
}

func (*SysTaskPlugin) RouterPath() string {
	return ""
}
