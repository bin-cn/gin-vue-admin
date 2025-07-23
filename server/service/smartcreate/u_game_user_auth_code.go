
package smartcreate

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
    smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
)

type GameUserAuthCodeService struct {}
// CreateGameUserAuthCode 创建用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService) CreateGameUserAuthCode(ctx context.Context, game_user_auth_code *smartcreate.GameUserAuthCode) (err error) {
	err = global.GVA_DB.Create(game_user_auth_code).Error
	return err
}

// DeleteGameUserAuthCode 删除用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService)DeleteGameUserAuthCode(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&smartcreate.GameUserAuthCode{},"id = ?",ID).Error
	return err
}

// DeleteGameUserAuthCodeByIds 批量删除用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService)DeleteGameUserAuthCodeByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]smartcreate.GameUserAuthCode{},"id in ?",IDs).Error
	return err
}

// UpdateGameUserAuthCode 更新用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService)UpdateGameUserAuthCode(ctx context.Context, game_user_auth_code smartcreate.GameUserAuthCode) (err error) {
	err = global.GVA_DB.Model(&smartcreate.GameUserAuthCode{}).Where("id = ?",game_user_auth_code.ID).Updates(&game_user_auth_code).Error
	return err
}

// GetGameUserAuthCode 根据ID获取用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService)GetGameUserAuthCode(ctx context.Context, ID string) (game_user_auth_code smartcreate.GameUserAuthCode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&game_user_auth_code).Error
	return
}
// GetGameUserAuthCodeInfoList 分页获取用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService)GetGameUserAuthCodeInfoList(ctx context.Context, info smartcreateReq.GameUserAuthCodeSearch) (list []smartcreate.GameUserAuthCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&smartcreate.GameUserAuthCode{})
    var game_user_auth_codes []smartcreate.GameUserAuthCode
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.LoginCode != nil && *info.LoginCode != "" {
        db = db.Where("login_code = ?", *info.LoginCode)
    }
    if info.MachineNoName != nil && *info.MachineNoName != "" {
        db = db.Where("machine_no_name = ?", *info.MachineNoName)
    }
    if info.GameServerName != nil && *info.GameServerName != "" {
        db = db.Where("game_server_name = ?", *info.GameServerName)
    }
	if info.StartGameServerId != nil && info.EndGameServerId != nil {
		db = db.Where("game_server_id BETWEEN ? AND ? ", *info.StartGameServerId, *info.EndGameServerId)
	}
    if info.RoleGameId != nil {
        db = db.Where("role_game_id = ?", *info.RoleGameId)
    }
    if info.IDName != nil && *info.IDName != "" {
        db = db.Where("i_d_name LIKE ?", "%"+ *info.IDName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["id"] = true
           orderMap["created_at"] = true
         	orderMap["machine_no_name"] = true
         	orderMap["assigner_name"] = true
         	orderMap["account"] = true
         	orderMap["game_server_name"] = true
         	orderMap["game_server_id"] = true
         	orderMap["role_game_id"] = true
         	orderMap["server_open_time"] = true
         	orderMap["enter_server_time"] = true
         	orderMap["account_status"] = true
         	orderMap["i_d_name"] = true
         	orderMap["i_d_card_number"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&game_user_auth_codes).Error
	return  game_user_auth_codes, total, err
}
func (game_user_auth_codeService *GameUserAuthCodeService)GetGameUserAuthCodePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
