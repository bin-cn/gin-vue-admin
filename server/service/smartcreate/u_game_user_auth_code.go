package smartcreate

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
)

type GameUserAuthCodeService struct{}

// CreateGameUserAuthCode 创建用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService) CreateGameUserAuthCode(ctx context.Context, game_user_auth_code *smartcreate.GameUserAuthCode) (err error) {
	// 检查role_game_id的唯一性
	if game_user_auth_code.RoleGameId != nil && *game_user_auth_code.RoleGameId != "" {
		var count int64
		err = global.GVA_DB.Model(&smartcreate.GameUserAuthCode{}).
			Where("role_game_id = ?", *game_user_auth_code.RoleGameId).
			Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("角色游戏ID '%s' 已存在", *game_user_auth_code.RoleGameId)
		}
	}

	err = global.GVA_DB.Create(game_user_auth_code).Error
	return err
}

// DeleteGameUserAuthCode 删除用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService) DeleteGameUserAuthCode(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&smartcreate.GameUserAuthCode{}, "id = ?", ID).Error
	return err
}

// DeleteGameUserAuthCodeByIds 批量删除用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService) DeleteGameUserAuthCodeByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]smartcreate.GameUserAuthCode{}, "id in ?", IDs).Error
	return err
}

// UpdateGameUserAuthCode 更新用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService) UpdateGameUserAuthCode(ctx context.Context, game_user_auth_code smartcreate.GameUserAuthCode) (err error) {
	// 检查role_game_id的唯一性（排除当前记录）
	if game_user_auth_code.RoleGameId != nil && *game_user_auth_code.RoleGameId != "" {
		var count int64
		err = global.GVA_DB.Model(&smartcreate.GameUserAuthCode{}).
			Where("role_game_id = ? AND id != ?", *game_user_auth_code.RoleGameId, game_user_auth_code.ID).
			Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("角色游戏ID '%s' 已存在", *game_user_auth_code.RoleGameId)
		}
	}

	err = global.GVA_DB.Model(&smartcreate.GameUserAuthCode{}).Where("id = ?", game_user_auth_code.ID).Updates(&game_user_auth_code).Error
	return err
}

// GetGameUserAuthCode 根据ID获取用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService) GetGameUserAuthCode(ctx context.Context, ID string) (game_user_auth_code smartcreate.GameUserAuthCode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&game_user_auth_code).Error
	return
}

// GetGameUserAuthCodeInfoList 分页获取用户授权码码登录信息记录
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService) GetGameUserAuthCodeInfoList(ctx context.Context, info smartcreateReq.GameUserAuthCodeSearch) (list []smartcreate.GameUserAuthCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&smartcreate.GameUserAuthCode{})
	var game_user_auth_codes []smartcreate.GameUserAuthCode
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", *info.UserId)
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
	if info.RoleGameId != nil && *info.RoleGameId != "" {
		db = db.Where("role_game_id = ?", *info.RoleGameId)
	}
	if info.IDName != nil && *info.IDName != "" {
		db = db.Where("i_d_name LIKE ?", "%"+*info.IDName+"%")
	}
	if info.MainServerZone != nil && *info.MainServerZone != "" {
		db = db.Where("main_server_zone = ?", *info.MainServerZone)
	}
	if info.AccountInternalId != nil && *info.AccountInternalId != "" {
		db = db.Where("account_internal_id = ?", *info.AccountInternalId)
	}
	err = db.Count(&total).Error
	if err != nil {
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
	orderMap["main_server_zone"] = true

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
	return game_user_auth_codes, total, err
}
func (game_user_auth_codeService *GameUserAuthCodeService) GetGameUserAuthCodePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetGameUserAuthCodeInfoListWithAuth 分页获取用户授权码码登录信息记录（带权限过滤）
// Author [yourname](https://github.com/yourname)
func (game_user_auth_codeService *GameUserAuthCodeService) GetGameUserAuthCodeInfoListWithAuth(ctx context.Context, info smartcreateReq.GameUserAuthCodeSearch, currentUserId uint, authorityId uint) (list []smartcreate.GameUserAuthCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&smartcreate.GameUserAuthCode{})
	var game_user_auth_codes []smartcreate.GameUserAuthCode

	// 如果角色ID为4，只显示当前用户的数据
	if authorityId == 4 {
		db = db.Where("user_id = ?", currentUserId)
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.UserId != nil {
		// 如果角色ID为4，不允许查询其他用户的数据
		if authorityId == 4 {
			if *info.UserId != int(currentUserId) {
				return []smartcreate.GameUserAuthCode{}, 0, nil // 返回空结果
			}
		}
		db = db.Where("user_id = ?", *info.UserId)
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
	if info.RoleGameId != nil && *info.RoleGameId != "" {
		db = db.Where("role_game_id = ?", *info.RoleGameId)
	}
	if info.IDName != nil && *info.IDName != "" {
		db = db.Where("i_d_name LIKE ?", "%"+*info.IDName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
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
	return game_user_auth_codes, total, err
}
