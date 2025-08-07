package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GameServerSearch struct {
	CreatedAtRange   []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	MainServerZoneId *string     `json:"mainServerZoneId" form:"mainServerZoneId"`
	ServerId         *string     `json:"serverId" form:"serverId"`
	ServerZoneId     *string     `json:"serverZoneId" form:"serverZoneId"`
	ServerName       *string     `json:"serverName" form:"serverName"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
