// 自动生成模板GameUser
package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户信息表 结构体  GameUser
type GameUser struct {
	global.GVA_MODEL
	NickName      *string `json:"nickName" form:"nickName" gorm:"comment:使用者的昵称;column:nick_name;" binding:"required"`                           //用户昵称
	LoginCode     *string `json:"loginCode" form:"loginCode" gorm:"uniqueIndex;comment:登录码,根据登录码获取用户名和密码;column:login_code;" binding:"required"` //登录码
	RoleGameName  *string `json:"roleGameName" form:"roleGameName" gorm:"comment:游戏角色名称;column:role_game_name;"`                                 //游戏角色名称
	RoleGameLevel *string `json:"roleGameLevel" form:"roleGameLevel" gorm:"column:role_game_level;"`                                             //游戏角色等级
	UserId        *int    `json:"userId" form:"userId" gorm:"index;comment:绑定用户的ID;column:user_id;" binding:"required"`                          //用户ID

	// 在结构体中新增如下字段
	GameServerName *string `json:"gameServerName" form:"gameServerName" gorm:"column:game_server_name;"`              //区服名称
	GameServerId   *int    `json:"gameServerId" form:"gameServerId" gorm:"column:game_server_id;" binding:"required"` //区服ID

	UnBoundIngotQuantity *int `json:"unBoundIngotQuantity" form:"unBoundIngotQuantity" gorm:"default:0;comment:未绑定元宝数量;column:un_bound_ingot_quantity;"` //未绑定元宝数量
	BoundIngotQuantity   *int `json:"boundIngotQuantity" form:"boundIngotQuantity" gorm:"default:0;comment:绑定元宝数量;column:bound_ingot_quantity;"`         //绑定元宝数量
	TotalIngotQuantity   *int `json:"totalIngotQuantity" form:"totalIngotQuantity" gorm:"default:0;comment:元宝总数;column:total_ingot_quantity;"`           //元宝总数
	UnBoundTalisman      *int `json:"unBoundTalisman" form:"unBoundTalisman" gorm:"default:0;comment:未绑定灵符;column:un_bound_talisman;"`                   //未绑定灵符
	BoundTalisman        *int `json:"boundTalisman" form:"boundTalisman" gorm:"default:0;comment:绑定灵符;column:bound_talisman;"`                           //绑定灵符
	TotalTalisman        *int `json:"totalTalisman" form:"totalTalisman" gorm:"default:0;comment:灵符总数;column:total_talisman;"`                           //灵符总数
	CreatedBy            uint `gorm:"column:created_by;comment:创建者"`
	UpdatedBy            uint `gorm:"column:updated_by;comment:更新者"`
	DeletedBy            uint `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 用户信息表 GameUser自定义表名 u_game_user
func (GameUser) TableName() string {
	return "u_game_user"
}
