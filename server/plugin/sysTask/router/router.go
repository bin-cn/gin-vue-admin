package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/api"
	"github.com/gin-gonic/gin"
)

type SysTaskRouter struct {
}

func (s *SysTaskRouter) InitSysTaskRouter(Router *gin.RouterGroup) {
	sysTaskRouter := Router.Group("sysTask").Use(middleware.OperationRecord())
	sysTaskLogRouter := Router.Group("sysTaskLog").Use(middleware.OperationRecord())
	sysTaskRouterWithoutRecord := Router.Group("sysTask")
	sysTaskLogRouterWithoutRecord := Router.Group("sysTaskLog")
	sysTaskApi := api.SysTaskGroupApp.SysTaskApi
	var sysTaskLogApi = api.SysTaskGroupApp.SysTaskLogApi
	{
		sysTaskRouter.POST("createSysTask", sysTaskApi.CreateSysTask)             // 新建sysTask表
		sysTaskRouter.DELETE("deleteSysTask", sysTaskApi.DeleteSysTask)           // 删除sysTask表
		sysTaskRouter.DELETE("deleteSysTaskByIds", sysTaskApi.DeleteSysTaskByIds) // 批量删除sysTask表
		sysTaskRouter.PUT("updateSysTask", sysTaskApi.UpdateSysTask)              // 更新sysTask表
	}
	{
		sysTaskRouterWithoutRecord.GET("findSysTask", sysTaskApi.FindSysTask)       // 根据ID获取sysTask表
		sysTaskRouterWithoutRecord.GET("getSysTaskList", sysTaskApi.GetSysTaskList) // 获取sysTask表列表
		sysTaskRouterWithoutRecord.GET("runOneSysTask", sysTaskApi.RunOneSysTask)   // 运行一次
		sysTaskRouterWithoutRecord.GET("enableSysTask", sysTaskApi.EnableSysTask)   // 激活任务
		sysTaskRouterWithoutRecord.GET("disableSysTask", sysTaskApi.DisableSysTask) // 暂停任务
		sysTaskRouterWithoutRecord.GET("removeSysTask", sysTaskApi.RemoveSysTask)   // 删除任务
	}
	{
		sysTaskLogRouter.POST("createSysTaskLog", sysTaskLogApi.CreateSysTaskLog)             // 新建sysTaskLog表
		sysTaskLogRouter.DELETE("deleteSysTaskLog", sysTaskLogApi.DeleteSysTaskLog)           // 删除sysTaskLog表
		sysTaskLogRouter.DELETE("deleteSysTaskLogByIds", sysTaskLogApi.DeleteSysTaskLogByIds) // 批量删除sysTaskLog表
		sysTaskLogRouter.PUT("updateSysTaskLog", sysTaskLogApi.UpdateSysTaskLog)              // 更新sysTaskLog表
	}
	{
		sysTaskLogRouterWithoutRecord.GET("findSysTaskLog", sysTaskLogApi.FindSysTaskLog)       // 根据ID获取sysTaskLog表
		sysTaskLogRouterWithoutRecord.GET("getSysTaskLogList", sysTaskLogApi.GetSysTaskLogList) // 获取sysTaskLog表列表
	}
}
