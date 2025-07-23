package smartcreate

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	GameUserRouter
	GameUserAuthCodeRouter
}

var (
	game_userApi           = api.ApiGroupApp.SmartcreateApiGroup.GameUserApi
	game_user_auth_codeApi = api.ApiGroupApp.SmartcreateApiGroup.GameUserAuthCodeApi
)
