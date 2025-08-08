package smartcreate

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	GameUserApi
	GameUserAuthCodeApi
	GameUserAuthCodeImportApi
	GameServerApi
	CookieDataApi
}

var (
	game_userService                 = service.ServiceGroupApp.SmartcreateServiceGroup.GameUserService
	game_user_auth_codeService       = service.ServiceGroupApp.SmartcreateServiceGroup.GameUserAuthCodeService
	game_user_auth_codeImportService = service.ServiceGroupApp.SmartcreateServiceGroup.GameUserAuthCodeImportService
	gamesrvService                   = service.ServiceGroupApp.SmartcreateServiceGroup.GameServerService
	wcService                        = service.ServiceGroupApp.SmartcreateServiceGroup.CookieDataService
)
