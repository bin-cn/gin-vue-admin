package smartcreate

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	GameUserApi
	GameUserAuthCodeApi
}

var (
	game_userService           = service.ServiceGroupApp.SmartcreateServiceGroup.GameUserService
	game_user_auth_codeService = service.ServiceGroupApp.SmartcreateServiceGroup.GameUserAuthCodeService
)
