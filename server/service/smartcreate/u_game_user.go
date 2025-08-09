package smartcreate

import (
	"context"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
	"gorm.io/gorm"
)

type GameUserService struct{}

// CreateGameUser 创建用户信息表记录
// Author [yourname](https://github.com/yourname)
func (game_userService *GameUserService) CreateGameUser(ctx context.Context, game_user *smartcreate.GameUser) (err error) {
	err = global.GVA_DB.Create(game_user).Error
	return err
}

// DeleteGameUser 删除用户信息表记录
// Author [yourname](https://github.com/yourname)
func (game_userService *GameUserService) DeleteGameUser(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&smartcreate.GameUser{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&smartcreate.GameUser{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteGameUserByIds 批量删除用户信息表记录
// Author [yourname](https://github.com/yourname)
func (game_userService *GameUserService) DeleteGameUserByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&smartcreate.GameUser{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&smartcreate.GameUser{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateGameUser 更新用户信息表记录
// Author [yourname](https://github.com/yourname)
func (game_userService *GameUserService) UpdateGameUser(ctx context.Context, game_user smartcreate.GameUser) (err error) {
	err = global.GVA_DB.Model(&smartcreate.GameUser{}).Where("id = ?", game_user.ID).Updates(&game_user).Error
	return err
}

// GetGameUser 根据ID获取用户信息表记录
// Author [yourname](https://github.com/yourname)
func (game_userService *GameUserService) GetGameUser(ctx context.Context, ID string) (game_user smartcreate.GameUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&game_user).Error
	return
}

// GetGameUserInfoList 分页获取用户信息表记录
// Author [yourname](https://github.com/yourname)
func (game_userService *GameUserService) GetGameUserInfoList(ctx context.Context, info smartcreateReq.GameUserSearch) (list []smartcreate.GameUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&smartcreate.GameUser{})
	var game_users []smartcreate.GameUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.NickName != nil && *info.NickName != "" {
		db = db.Where("nick_name = ?", *info.NickName)
	}
	if info.LoginCode != nil && *info.LoginCode != "" {
		db = db.Where("login_code = ?", *info.LoginCode)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", *info.UserId)
	}
	if info.StartUnBoundIngotQuantity != nil && info.EndUnBoundIngotQuantity != nil {
		db = db.Where("un_bound_ingot_quantity BETWEEN ? AND ? ", *info.StartUnBoundIngotQuantity, *info.EndUnBoundIngotQuantity)
	}
	if info.StartBoundIngotQuantity != nil && info.EndBoundIngotQuantity != nil {
		db = db.Where("bound_ingot_quantity BETWEEN ? AND ? ", *info.StartBoundIngotQuantity, *info.EndBoundIngotQuantity)
	}
	if info.StartTotalIngotQuantity != nil && info.EndTotalIngotQuantity != nil {
		db = db.Where("total_ingot_quantity BETWEEN ? AND ? ", *info.StartTotalIngotQuantity, *info.EndTotalIngotQuantity)
	}
	if info.StartUnBoundTalisman != nil && info.EndUnBoundTalisman != nil {
		db = db.Where("un_bound_talisman BETWEEN ? AND ? ", *info.StartUnBoundTalisman, *info.EndUnBoundTalisman)
	}
	if info.StartBoundTalisman != nil && info.EndBoundTalisman != nil {
		db = db.Where("bound_talisman BETWEEN ? AND ? ", *info.StartBoundTalisman, *info.EndBoundTalisman)
	}
	if info.StartTotalTalisman != nil && info.EndTotalTalisman != nil {
		db = db.Where("total_talisman BETWEEN ? AND ? ", *info.StartTotalTalisman, *info.EndTotalTalisman)
	}

	// GetGameUserInfoList 新增搜索语句

	if info.GameServerName != nil && *info.GameServerName != "" {
		db = db.Where("game_server_name LIKE ?", "%"+*info.GameServerName+"%")
	}
	if info.StartGameServerId != nil && info.EndGameServerId != nil {
		db = db.Where("game_server_id BETWEEN ? AND ? ", *info.StartGameServerId, *info.EndGameServerId)
	}

	if info.RoleOnlineStatus != nil && *info.RoleOnlineStatus != "" {
		db = db.Where("role_online_status = ?", *info.RoleOnlineStatus)
	}
	if info.ScriptOnlineStatus != nil && *info.ScriptOnlineStatus != "" {
		db = db.Where("script_online_status = ?", *info.ScriptOnlineStatus)
	}
	if info.BannedStatus != nil && *info.BannedStatus != "" {
		db = db.Where("banned_status = ?", *info.BannedStatus)
	}
	if info.StartTalismanDiff != nil && info.EndTalismanDiff != nil {
		db = db.Where("talisman_diff BETWEEN ? AND ? ", *info.StartTalismanDiff, *info.EndTalismanDiff)
	}
	if info.StartIngotDiff != nil && info.EndIngotDiff != nil {
		db = db.Where("ingot_diff BETWEEN ? AND ? ", *info.StartIngotDiff, *info.EndIngotDiff)
	}
	if len(info.LastSyncQueryTimeRange) == 2 {
		db = db.Where("last_sync_query_time BETWEEN ? AND ? ", info.LastSyncQueryTimeRange[0], info.LastSyncQueryTimeRange[1])
	}
	if len(info.LastSyncUpdateTimeRange) == 2 {
		db = db.Where("last_sync_update_time BETWEEN ? AND ? ", info.LastSyncUpdateTimeRange[0], info.LastSyncUpdateTimeRange[1])
	}
	if len(info.ScriptLastOnlineTimeRange) == 2 {
		db = db.Where("script_last_online_time BETWEEN ? AND ? ", info.ScriptLastOnlineTimeRange[0], info.ScriptLastOnlineTimeRange[1])
	}
	if len(info.RoleLastOnlineTimeRange) == 2 {
		db = db.Where("role_last_online_time BETWEEN ? AND ? ", info.RoleLastOnlineTimeRange[0], info.RoleLastOnlineTimeRange[1])
	}
	if info.RoleGameId != nil && *info.RoleGameId != "" {
		db = db.Where("role_game_id = ?", *info.RoleGameId)
	}
	if info.ServerZoneId != nil && *info.ServerZoneId != "" {
		db = db.Where("server_zone_id = ?", *info.ServerZoneId)
	}
	if len(info.LastIngotTradeTimeRange) == 2 {
		db = db.Where("last_ingot_trade_time BETWEEN ? AND ? ", info.LastIngotTradeTimeRange[0], info.LastIngotTradeTimeRange[1])
	}
	if len(info.LastTalismanTradeTimeRange) == 2 {
		db = db.Where("last_talisman_trade_time BETWEEN ? AND ? ", info.LastTalismanTradeTimeRange[0], info.LastTalismanTradeTimeRange[1])
	}
	if info.OnlineTalismanTotal != nil {
		db = db.Where("online_talisman_total > ?", *info.OnlineTalismanTotal)
	}
	if info.OnlineIngotTotal != nil {
		db = db.Where("online_ingot_total > ?", *info.OnlineIngotTotal)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["nick_name"] = true
	orderMap["role_game_level"] = true
	orderMap["un_bound_ingot_quantity"] = true
	orderMap["bound_ingot_quantity"] = true
	orderMap["total_ingot_quantity"] = true
	orderMap["un_bound_talisman"] = true
	orderMap["bound_talisman"] = true
	orderMap["total_talisman"] = true
	// GetGameUserInfoList 新增排序语句 请自行在搜索语句中添加orderMap内容
	orderMap["game_server_name"] = true
	orderMap["game_server_id"] = true
	orderMap["role_online_status"] = true
	orderMap["script_online_status"] = true
	orderMap["banned_status"] = true
	orderMap["talisman_diff"] = true
	orderMap["ingot_diff"] = true
	orderMap["last_sync_query_time"] = true
	orderMap["last_sync_update_time"] = true
	orderMap["script_last_online_time"] = true
	orderMap["role_last_online_time"] = true
	orderMap["server_zone_id"] = true
	orderMap["last_ingot_trade_time"] = true
	orderMap["last_talisman_trade_time"] = true
	orderMap["online_talisman_total"] = true
	orderMap["online_ingot_total"] = true
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

	err = db.Find(&game_users).Error
	return game_users, total, err
}
func (game_userService *GameUserService) GetGameUserDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	nickName := make([]map[string]any, 0)

	global.GVA_DB.Table("sys_users").Where("deleted_at IS NULL").Select("nick_name as label,nick_name as value").Scan(&nickName)
	res["nickName"] = nickName
	userId := make([]map[string]any, 0)

	global.GVA_DB.Table("sys_users").Where("deleted_at IS NULL").Select("id as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}
func (game_userService *GameUserService) GetGameUserPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// SubmitGoldCoinInfo 提交用户资产信息
// Author [yourname](https://github.com/yourname)
func (game_userService *GameUserService) SubmitGoldCoinInfo(ctx context.Context, game_user smartcreateReq.SubmitGoldCoinInfo) (err error) {
	var count int64
	var existingUser *smartcreate.GameUser
	var userId uint = 0

	// 1.查询数据库中是否有这个授权码,如果有则更新.
	//  这里实际上应该还要查询一个UserId,这里暂时先不处理,先把导入处理完毕以后再加入.
	err = global.GVA_DB.WithContext(ctx).Model(&smartcreate.GameUser{}).
		Where("login_code = ?", game_user.LoginCode).
		First(&existingUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("未能找到该授权码对应用户: %w", err)
	}

	if err != nil && existingUser != nil {
		userId = uint(*existingUser.UserId)
		if userId != game_user.UpdatedBy {
			fmt.Print("这里临时输出日志,后需需要判断该授权码是否属于当前用户..")
		}
	}

	// 若存在，手动指定需要更新的字段
	if err == nil && existingUser != nil {
		// 仅更新需要的字段（示例：未绑定元宝、绑定元宝、总元宝等）
		updateData := map[string]interface{}{
			"bound_ingot_quantity":    *game_user.TotalIngotQuantity - *game_user.UnBoundIngotQuantity,
			"bound_talisman":          *game_user.TotalTalisman - *game_user.UnBoundTalisman,
			"total_ingot_quantity":    *game_user.TotalIngotQuantity,
			"total_talisman":          *game_user.TotalTalisman,
			"un_bound_ingot_quantity": *game_user.UnBoundIngotQuantity,
			"un_bound_talisman":       *game_user.UnBoundTalisman,

			// 可根据需求添加其他需要更新的字段
		}

		err = global.GVA_DB.WithContext(ctx).Model(&smartcreate.GameUser{}).
			Where("id = ?", existingUser.ID).
			Updates(updateData).Error // 仅更新指定字段
		if err != nil {
			return fmt.Errorf("更新用户信息失败: %w", err)
		}
		return nil
	}

	// 2. 查询授权码信息
	err = global.GVA_DB.WithContext(ctx).Model(&smartcreate.GameUserAuthCode{}).Where("login_code = ?", game_user.LoginCode).
		Count(&count).Error
	if err != nil {
		return fmt.Errorf("授权码不存在.")
	}
	if count == 0 {
		return errors.New("login_code不存在，无法创建用户信息")
	}

	// 如果没有则从授权信息中获取这个授权信息插入进去,更新
	// err = game_userService.CreateGameUser(ctx, gameUser)
	// if err != nil {
	// 	return fmt.Errorf("插入用户信息失败: %w", err)
	// }
	// // 增加日志记录.

	// // 请在这里实现自己的业务逻辑
	// db := global.GVA_DB.Model(&smartcreate.GameUser{})
	// return db.Error
	return nil
}
