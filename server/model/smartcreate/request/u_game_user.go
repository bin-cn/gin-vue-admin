package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SubmitGoldCoinInfo struct {
	LoginCode            *string `json:"loginCode" form:"loginCode" gorm:"uniqueIndex;comment:登录码,根据登录码获取用户名和密码;column:login_code;" binding:"required"`     //登录码
	UserId               *int    `json:"userId" form:"userId" gorm:"index;comment:绑定用户的ID;column:user_id;"`                                                 //用户ID
	UnBoundIngotQuantity *int    `json:"unBoundIngotQuantity" form:"unBoundIngotQuantity" gorm:"default:0;comment:未绑定元宝数量;column:un_bound_ingot_quantity;"` //未绑定元宝数量
	TotalIngotQuantity   *int    `json:"totalIngotQuantity" form:"totalIngotQuantity" gorm:"default:0;comment:元宝总数;column:total_ingot_quantity;"`           //元宝总数
	UnBoundTalisman      *int    `json:"unBoundTalisman" form:"unBoundTalisman" gorm:"default:0;comment:未绑定灵符;column:un_bound_talisman;"`                   //未绑定灵符
	TotalTalisman        *int    `json:"totalTalisman" form:"totalTalisman" gorm:"default:0;comment:灵符总数;column:total_talisman;"`                           //灵符总数
	UpdatedBy            uint    `gorm:"column:updated_by;comment:更新者"`
}

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
