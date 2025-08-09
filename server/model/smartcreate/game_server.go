// 自动生成模板GameServer
package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 游戏服务器列表 结构体  GameServer
type GameServer struct {
	global.GVA_MODEL
	MainServerZoneId *string `json:"mainServerZoneId" form:"mainServerZoneId" gorm:"comment:主区服ID;column:main_server_zone_id;"` //主区服ID
	MainServerId     *string `json:"mainServerId" form:"mainServerId" gorm:"comment:主服务器ID;column:main_server_id;"`             //主服务器ID
	ServerId         *string `json:"serverId" form:"serverId" gorm:"comment:服务器ID;column:server_id;"`                           //服务器ID
	ServerZoneId     *string `json:"serverZoneId" form:"serverZoneId" gorm:"comment:区服ID;column:server_zone_id;"`               //区服ID
	ServerName       *string `json:"serverName" form:"serverName" gorm:"comment:区服名字;column:server_name;"`                      //区服名字
}

// TableName 游戏服务器列表 GameServer自定义表名 game_server
func (GameServer) TableName() string {
	return "game_server"
}
