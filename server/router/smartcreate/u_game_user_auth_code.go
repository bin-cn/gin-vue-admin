package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GameUserAuthCodeRouter struct {}

// InitGameUserAuthCodeRouter 初始化 用户授权码码登录信息 路由信息
func (s *GameUserAuthCodeRouter) InitGameUserAuthCodeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	game_user_auth_codeRouter := Router.Group("game_user_auth_code").Use(middleware.OperationRecord())
	game_user_auth_codeRouterWithoutRecord := Router.Group("game_user_auth_code")
	game_user_auth_codeRouterWithoutAuth := PublicRouter.Group("game_user_auth_code")
	{
		game_user_auth_codeRouter.POST("createGameUserAuthCode", game_user_auth_codeApi.CreateGameUserAuthCode)   // 新建用户授权码码登录信息
		game_user_auth_codeRouter.DELETE("deleteGameUserAuthCode", game_user_auth_codeApi.DeleteGameUserAuthCode) // 删除用户授权码码登录信息
		game_user_auth_codeRouter.DELETE("deleteGameUserAuthCodeByIds", game_user_auth_codeApi.DeleteGameUserAuthCodeByIds) // 批量删除用户授权码码登录信息
		game_user_auth_codeRouter.PUT("updateGameUserAuthCode", game_user_auth_codeApi.UpdateGameUserAuthCode)    // 更新用户授权码码登录信息
	}
	{
		game_user_auth_codeRouterWithoutRecord.GET("findGameUserAuthCode", game_user_auth_codeApi.FindGameUserAuthCode)        // 根据ID获取用户授权码码登录信息
		game_user_auth_codeRouterWithoutRecord.GET("getGameUserAuthCodeList", game_user_auth_codeApi.GetGameUserAuthCodeList)  // 获取用户授权码码登录信息列表
	}
	{
	    game_user_auth_codeRouterWithoutAuth.GET("getGameUserAuthCodePublic", game_user_auth_codeApi.GetGameUserAuthCodePublic)  // 用户授权码码登录信息开放接口
	}
}
