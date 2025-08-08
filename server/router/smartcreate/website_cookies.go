package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CookieDataRouter struct {}

// InitCookieDataRouter 初始化 网站cookie 路由信息
func (s *CookieDataRouter) InitCookieDataRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcRouter := Router.Group("wc").Use(middleware.OperationRecord())
	wcRouterWithoutRecord := Router.Group("wc")
	wcRouterWithoutAuth := PublicRouter.Group("wc")
	{
		wcRouter.POST("createCookieData", wcApi.CreateCookieData)   // 新建网站cookie
		wcRouter.DELETE("deleteCookieData", wcApi.DeleteCookieData) // 删除网站cookie
		wcRouter.DELETE("deleteCookieDataByIds", wcApi.DeleteCookieDataByIds) // 批量删除网站cookie
		wcRouter.PUT("updateCookieData", wcApi.UpdateCookieData)    // 更新网站cookie
	}
	{
		wcRouterWithoutRecord.GET("findCookieData", wcApi.FindCookieData)        // 根据ID获取网站cookie
		wcRouterWithoutRecord.GET("getCookieDataList", wcApi.GetCookieDataList)  // 获取网站cookie列表
	}
	{
	    wcRouterWithoutAuth.GET("getCookieDataPublic", wcApi.GetCookieDataPublic)  // 网站cookie开放接口
	}
}
