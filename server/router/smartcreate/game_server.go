package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GameServerRouter struct {}

// InitGameServerRouter 初始化 游戏服务器列表 路由信息
func (s *GameServerRouter) InitGameServerRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	gamesrvRouter := Router.Group("gamesrv").Use(middleware.OperationRecord())
	gamesrvRouterWithoutRecord := Router.Group("gamesrv")
	gamesrvRouterWithoutAuth := PublicRouter.Group("gamesrv")
	{
		gamesrvRouter.POST("createGameServer", gamesrvApi.CreateGameServer)   // 新建游戏服务器列表
		gamesrvRouter.DELETE("deleteGameServer", gamesrvApi.DeleteGameServer) // 删除游戏服务器列表
		gamesrvRouter.DELETE("deleteGameServerByIds", gamesrvApi.DeleteGameServerByIds) // 批量删除游戏服务器列表
		gamesrvRouter.PUT("updateGameServer", gamesrvApi.UpdateGameServer)    // 更新游戏服务器列表
	}
	{
		gamesrvRouterWithoutRecord.GET("findGameServer", gamesrvApi.FindGameServer)        // 根据ID获取游戏服务器列表
		gamesrvRouterWithoutRecord.GET("getGameServerList", gamesrvApi.GetGameServerList)  // 获取游戏服务器列表列表
	}
	{
	    gamesrvRouterWithoutAuth.GET("getGameServerPublic", gamesrvApi.GetGameServerPublic)  // 游戏服务器列表开放接口
	}
}
