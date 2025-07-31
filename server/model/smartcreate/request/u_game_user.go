package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GameUserSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	NickName       *string     `json:"nickName" form:"nickName"`
	LoginCode      *string     `json:"loginCode" form:"loginCode"`
	UserId         *int        `json:"userId" form:"userId"`

	// 在结构体中新增如下字段
	GameServerName            *string `json:"gameServerName" form:"gameServerName"`
	StartGameServerId         *int    `json:"startGameServerId" form:"startGameServerId"`
	EndGameServerId           *int    `json:"endGameServerId" form:"endGameServerId"`
	StartUnBoundIngotQuantity *int    `json:"startUnBoundIngotQuantity" form:"startUnBoundIngotQuantity"`
	EndUnBoundIngotQuantity   *int    `json:"endUnBoundIngotQuantity" form:"endUnBoundIngotQuantity"`
	StartBoundIngotQuantity   *int    `json:"startBoundIngotQuantity" form:"startBoundIngotQuantity"`
	EndBoundIngotQuantity     *int    `json:"endBoundIngotQuantity" form:"endBoundIngotQuantity"`
	StartTotalIngotQuantity   *int    `json:"startTotalIngotQuantity" form:"startTotalIngotQuantity"`
	EndTotalIngotQuantity     *int    `json:"endTotalIngotQuantity" form:"endTotalIngotQuantity"`
	StartUnBoundTalisman      *int    `json:"startUnBoundTalisman" form:"startUnBoundTalisman"`
	EndUnBoundTalisman        *int    `json:"endUnBoundTalisman" form:"endUnBoundTalisman"`
	StartBoundTalisman        *int    `json:"startBoundTalisman" form:"startBoundTalisman"`
	EndBoundTalisman          *int    `json:"endBoundTalisman" form:"endBoundTalisman"`
	StartTotalTalisman        *int    `json:"startTotalTalisman" form:"startTotalTalisman"`
	EndTotalTalisman          *int    `json:"endTotalTalisman" form:"endTotalTalisman"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
