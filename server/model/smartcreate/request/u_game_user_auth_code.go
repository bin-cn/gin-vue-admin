package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GameUserAuthCodeSearch struct {
	CreatedAtRange    []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	LoginCode         *string     `json:"loginCode" form:"loginCode"`
	MachineNoName     *string     `json:"machineNoName" form:"machineNoName"`
	GameServerName    *string     `json:"gameServerName" form:"gameServerName"`
	StartGameServerId *int        `json:"startGameServerId" form:"startGameServerId"`
	EndGameServerId   *int        `json:"endGameServerId" form:"endGameServerId"`
	RoleGameId        *string     `json:"roleGameId" form:"roleGameId"`
	IDName            *string     `json:"iDName" form:"iDName"`
	MainServerZone    *string     `json:"mainServerZone" form:"mainServerZone"`
	AccountInternalId *string     `json:"accountInternalId" form:"accountInternalId"`
	request.PageInfo
	Sort   string `json:"sort" form:"sort"`
	Order  string `json:"order" form:"order"`
	UserId *int   `json:"userId" form:"userId"`
}
