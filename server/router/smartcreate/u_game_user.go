package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GameUserRouter struct{}

func (s *GameUserRouter) InitGameUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	game_userRouter := Router.Group("game_user").Use(middleware.OperationRecord())
	game_userRouterWithoutRecord := Router.Group("game_user")
	game_userRouterWithoutAuth := PublicRouter.Group("game_user")
	{
		game_userRouter.POST("createGameUser", game_userApi.CreateGameUser)
		game_userRouter.DELETE("deleteGameUser", game_userApi.DeleteGameUser)
		game_userRouter.DELETE("deleteGameUserByIds", game_userApi.DeleteGameUserByIds)
		game_userRouter.PUT("updateGameUser", game_userApi.UpdateGameUser)
		game_userRouter.POST("submit_gold_coin_info", game_userApi.SubmitGoldCoinInfo)
	}
	{
		game_userRouterWithoutRecord.GET("findGameUser", game_userApi.FindGameUser)
		game_userRouterWithoutRecord.GET("getGameUserList", game_userApi.GetGameUserList)
	}
	{
		game_userRouterWithoutAuth.GET("getGameUserDataSource", game_userApi.GetGameUserDataSource)
		game_userRouterWithoutAuth.GET("getGameUserPublic", game_userApi.GetGameUserPublic)
	}
}
