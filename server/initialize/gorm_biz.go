package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(smartcreate.GameUser{}, smartcreate.GameUserAuthCode{}, smartcreate.GameServer{}, smartcreate.CookieData{})
	if err != nil {
		return err
	}
	return nil
}
