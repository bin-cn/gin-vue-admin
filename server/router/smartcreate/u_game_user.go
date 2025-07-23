package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GameUserRouter struct{}

// InitGameUserRouter 初始化 用户信息表 路由信息
func (s *GameUserRouter) InitGameUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	game_userRouter := Router.Group("game_user").Use(middleware.OperationRecord())
	game_userRouterWithoutRecord := Router.Group("game_user")
	game_userRouterWithoutAuth := PublicRouter.Group("game_user")
	{
		game_userRouter.POST("createGameUser", game_userApi.CreateGameUser)             // 新建用户信息表
		game_userRouter.DELETE("deleteGameUser", game_userApi.DeleteGameUser)           // 删除用户信息表
		game_userRouter.DELETE("deleteGameUserByIds", game_userApi.DeleteGameUserByIds) // 批量删除用户信息表
		game_userRouter.PUT("updateGameUser", game_userApi.UpdateGameUser)              // 更新用户信息表
	}
	{
		game_userRouterWithoutRecord.GET("findGameUser", game_userApi.FindGameUser)       // 根据ID获取用户信息表
		game_userRouterWithoutRecord.GET("getGameUserList", game_userApi.GetGameUserList) // 获取用户信息表列表
	}
	{
		game_userRouterWithoutAuth.GET("getGameUserDataSource", game_userApi.GetGameUserDataSource) // 获取用户信息表数据源
		game_userRouterWithoutAuth.GET("getGameUserPublic", game_userApi.GetGameUserPublic)         // 用户信息表开放接口
	}
}
