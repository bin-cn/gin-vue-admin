package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DailyRevenueRecordRouter struct {}

// InitDailyRevenueRecordRouter 初始化 日收入统计 路由信息
func (s *DailyRevenueRecordRouter) InitDailyRevenueRecordRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	drRouter := Router.Group("dr").Use(middleware.OperationRecord())
	drRouterWithoutRecord := Router.Group("dr")
	drRouterWithoutAuth := PublicRouter.Group("dr")
	{
		drRouter.POST("createDailyRevenueRecord", drApi.CreateDailyRevenueRecord)   // 新建日收入统计
		drRouter.DELETE("deleteDailyRevenueRecord", drApi.DeleteDailyRevenueRecord) // 删除日收入统计
		drRouter.DELETE("deleteDailyRevenueRecordByIds", drApi.DeleteDailyRevenueRecordByIds) // 批量删除日收入统计
		drRouter.PUT("updateDailyRevenueRecord", drApi.UpdateDailyRevenueRecord)    // 更新日收入统计
	}
	{
		drRouterWithoutRecord.GET("findDailyRevenueRecord", drApi.FindDailyRevenueRecord)        // 根据ID获取日收入统计
		drRouterWithoutRecord.GET("getDailyRevenueRecordList", drApi.GetDailyRevenueRecordList)  // 获取日收入统计列表
	}
	{
	    drRouterWithoutAuth.GET("getDailyRevenueRecordPublic", drApi.GetDailyRevenueRecordPublic)  // 日收入统计开放接口
	}
}
