
// 自动生成模板GameUserAuthCode
package smartcreate
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户授权码码登录信息 结构体  GameUserAuthCode
type GameUserAuthCode struct {
    global.GVA_MODEL
  LoginCode  *string `json:"loginCode" form:"loginCode" gorm:"uniqueIndex;comment:登录码;column:login_code;" binding:"required"`  //登录码
  MachineNoName  *string `json:"machineNoName" form:"machineNoName" gorm:"comment:机器编号;column:machine_no_name;" binding:"required"`  //机器编号
  AssignerName  *string `json:"assignerName" form:"assignerName" gorm:"comment:使用人;column:assigner_name;" binding:"required"`  //使用人
  Account  *string `json:"account" form:"account" gorm:"index;comment:账号;column:account;" binding:"required"`  //账号
  Password  *string `json:"password" form:"password" gorm:"comment:密码;column:password;" binding:"required"`  //密码
  GameServerName  *string `json:"gameServerName" form:"gameServerName" gorm:"column:game_server_name;" binding:"required"`  //区服
  GameServerId  *int `json:"gameServerId" form:"gameServerId" gorm:"index;comment:区服id;column:game_server_id;" binding:"required"`  //区服id
  RoleGameName  *string `json:"roleGameName" form:"roleGameName" gorm:"comment:游戏角色名字;column:role_game_name;" binding:"required"`  //游戏角色名字
  RoleGameId  *int `json:"roleGameId" form:"roleGameId" gorm:"comment:游戏角色ID;column:role_game_id;" binding:"required"`  //游戏角色ID
  ServerOpenTime  *string `json:"serverOpenTime" form:"serverOpenTime" gorm:"comment:开区时间;column:server_open_time;"`  //开区时间
  EnterServerTime  *string `json:"enterServerTime" form:"enterServerTime" gorm:"comment:进区时间;column:enter_server_time;"`  //进区时间
  AccountStatus  *string `json:"accountStatus" form:"accountStatus" gorm:"default:InUse;comment:账号状态;column:account_status;"`  //账号状态
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;"`  //备注
  IDName  *string `json:"iDName" form:"iDName" gorm:"comment:ID名字;column:i_d_name;" binding:"required"`  //ID名字
  IDCardNumber  *string `json:"iDCardNumber" form:"iDCardNumber" gorm:"comment:身份证号码;column:i_d_card_number;" binding:"required"`  //身份证号码
}


// TableName 用户授权码码登录信息 GameUserAuthCode自定义表名 game_user_auth_code
func (GameUserAuthCode) TableName() string {
    return "game_user_auth_code"
}





