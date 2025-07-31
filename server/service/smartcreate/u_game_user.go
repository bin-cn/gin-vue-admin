package smartcreate

import (
	"context"

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
