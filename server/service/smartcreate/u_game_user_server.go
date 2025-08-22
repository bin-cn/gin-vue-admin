package smartcreate

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	sc "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	"github.com/flipped-aurora/gin-vue-admin/server/service/gameManage"

	// service "github.com/flipped-aurora/gin-vue-admin/server/service/smartcreate"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MyError struct {
	Code int
	Msg  string
}

// 实现 error 接口
func (e *MyError) Error() string {
	return e.Msg
}

func VerifyItemUpdateLogsData(arg interface{}) (*gameManage.GameManager, error) {

	global.GVA_LOG.Info("获取物品更新日志任务开始")

	id := arg.(string)
	var cookieData *string = nil
	// 直接创建实例（不需要New方法）
	service := CookieDataService{}

	data, err := service.GetCookieData(nil, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没查到数据
			return nil, fmt.Errorf("未找到记录对应的cookie记录数据")
		}
		// 其他数据库错误
		return nil, fmt.Errorf("查询错误: %v", err)
	}

	if *data.Status == 0 {
		return nil, fmt.Errorf("cookie记录数据未启用")
	}

	cookieData = data.CookieValue
	gm := gameManage.NewGameManager(*cookieData)
	return gm, nil
}

// 包级变量：为每个 mainServerZoneStr 维护一把锁
var zoneLocks sync.Map // key:
// 自定义错误
var ErrZoneLocked = errors.New("该区服正在查询中.请勿重复查询.请等待20秒.刷新页面即可")

func Update_player_info_by_serverid_lock(gm *gameManage.GameManager, Action_type int, mainServerZoneStr string, time_type int) error {
	var err error
	// 1. 取出或创建这把锁
	lockIface, _ := zoneLocks.LoadOrStore(mainServerZoneStr, &sync.Mutex{})
	lock := lockIface.(*sync.Mutex)

	// 2. **非阻塞**加锁：如果拿不到立即返回提示
	if !lock.TryLock() {
		return ErrZoneLocked
	}
	// 3. 拿到锁后，确保最终解锁
	defer lock.Unlock()

	var my_err error

	err = update_player_info_by_serverid(gm, Action_type, UpdateTypeIngot, mainServerZoneStr, time_type)
	if err != nil {
		my_err = &MyError{Code: 99, Msg: "更新玩家[元宝]信息失败" + err.Error()}
		return my_err
	}

	// err = update_player_info_by_serverid(gm, Action_type, UpdateTypeTalisman, mainServerZoneStr, time_type)
	// if err != nil {
	// 	my_err = &MyError{Code: 99, Msg: "更新玩家[灵符]信息失败" + err.Error()}
	// 	return my_err
	// }

	return err

}

func update_player_info_by_serverid(gm *gameManage.GameManager, Action_type int, update_type int, mainServerZoneStr string, time_type int) error {

	//1.2查询出u_game_user表中的所有数据.
	var gameUsers []smartcreate.GameUser
	var game_user_auth_codes []smartcreate.GameUserAuthCode
	startTime, endTime := gameManage.BuildTime(int(time_type))

	var err error = nil
	var playerActions []gameManage.PlayerAction

	var item = ""

	switch update_type {
	case UpdateTypeTalisman:
		item = "2:灵符"
	case UpdateTypeIngot:
		item = "20:元宝"
	}

	playerActions, err = gm.GetPlayerActions(startTime, endTime, item, mainServerZoneStr)
	if err != nil {
		return fmt.Errorf("获取玩家操作失败: %v", err)
	}

	if len(playerActions) == 0 {
		return fmt.Errorf("获取玩家操作失败: 该区服没有玩家操作数据")
	}

	//1.1 查询game_user_auth_code表中的所有数据.
	db := global.GVA_DB.Model(&smartcreate.GameUserAuthCode{})
	err = db.Find(&game_user_auth_codes).Error
	if err != nil {
		return fmt.Errorf("查询 [授权表信息表]中的所有数据失败: %v", err)
	}

	db = global.GVA_DB.Model(&smartcreate.GameUser{})
	err = db.Find(&gameUsers).Error
	if err != nil {
		return fmt.Errorf("查询[用户游戏表]表中的所有数据失败: %v", err)
	}
	diff_update_data(playerActions, game_user_auth_codes, gameUsers, update_type, Action_type)
	return nil
}

// 定义枚举类型
const (
	// 更新类型
	UpdateTypeTalisman = 2  // 非绑定灵符
	UpdateTypeIngot    = 20 // 非绑定元宝

	// 操作类型
	ActionTypeScheduled = 1 // 定时任务更新
	ActionTypeManual    = 2 // 用户手动更新
)

// 修改diff_update_data方法中的字段处理
func diff_update_data(playerActions []gameManage.PlayerAction, game_user_auth_codes []sc.GameUserAuthCode, gameUsers []sc.GameUser, update_type int, Action_type int) error {
	// 返回更新后的gameUsers
	// 2.字段映射关系 ：这两张表的数据，字段一样的进行更新即可，不一样的不用处理，
	// 3.更新优先级 ： 我都说了以GameUser表中的数据为准，他的优先级最高，
	// 4. 数据不一致则输出错误信息，略过即可，
	var installUsers []sc.GameUser

	// 1. 对playerActions按PlayerId分组，取最新时间的数据
	latestActions := make(map[string]gameManage.PlayerAction)
	for _, action := range playerActions {
		if action.PlayerId == "" {
			continue // 跳过空的PlayerId
		}
		if existing, exists := latestActions[action.PlayerId]; !exists || action.Time > existing.Time {
			latestActions[action.PlayerId] = action
		}
	}

	// 2. 构建快速查找map
	gameUserMap := make(map[string]*sc.GameUser)
	for i := range gameUsers {
		if gameUsers[i].RoleGameId != nil {
			gameUserMap[*gameUsers[i].RoleGameId] = &gameUsers[i]
		}
	}

	authCodeMap := make(map[string]*sc.GameUserAuthCode)
	for i := range game_user_auth_codes {
		if game_user_auth_codes[i].RoleGameId != nil {
			authCodeMap[*game_user_auth_codes[i].RoleGameId] = &game_user_auth_codes[i]
		}
	}

	var updatedUsers []sc.GameUser // 收集更新的用户
	now := time.Now()

	// 3. 处理每个玩家的数据更新
	for playerId, action := range latestActions {

		// 4. 根据Action_type更新对应字段
		currentCountStr := action.CurrentCount

		currentCountInt64, err := strconv.ParseInt(currentCountStr, 10, 32)
		if err != nil {
			continue
		}
		currentCount := int(currentCountInt64)

		// 检查物品ID是否匹配更新类型
		if update_type == UpdateTypeTalisman && action.ItemId != "2" {
			continue
		}
		if update_type == UpdateTypeIngot && action.ItemId != "20" {
			continue
		}

		// 查找GameUser数据
		gameUser, exists := gameUserMap[playerId]
		if !exists {
			// GameUser中不存在，从GameUserAuthCode查找
			authCode, authExists := authCodeMap[playerId]
			if !authExists {
				continue
			}

			// 创建新的GameUser记录，从GameUserAuthCode复制相关字段
			mainServerZone := authCode.MainServerZone
			newUser := sc.GameUser{
				RoleGameId:           authCode.RoleGameId,
				NickName:             authCode.AssignerName, // 使用游戏角色名称作为昵称
				LoginCode:            authCode.LoginCode,
				RoleGameName:         authCode.RoleGameName,
				GameServerName:       authCode.GameServerName,
				GameServerId:         authCode.GameServerId, // 区服ID
				ServerZoneId:         mainServerZone,        // 主区服ID
				UnBoundIngotQuantity: new(int),              // 默认值0
				BoundIngotQuantity:   new(int),              // 默认值0
				TotalIngotQuantity:   new(int),              // 默认值0
				UnBoundTalisman:      new(int),              // 默认值0
				BoundTalisman:        new(int),              // 默认值0
				TotalTalisman:        new(int),              // 默认值0
				RoleOnlineStatus:     new(string),
				ScriptOnlineStatus:   new(string),
				BannedStatus:         new(string),
				UserId:               authCode.UserId,
			}
			*newUser.RoleOnlineStatus = "Offline"
			*newUser.ScriptOnlineStatus = "Offline"
			*newUser.BannedStatus = "0"
			newUser.LastOnlineQueryTime = &now
			newUser.LastSyncUpdateTime = &now
			newUser.LastSyncUpdateTime = &now

			if action.ItemId == "2" {
				// 2 是灵符
				newUser.UnBoundTalisman = &currentCount
				newUser.OnlineTalismanTotal = &currentCount
			} else if action.ItemId == "20" {
				//20 是元宝
				newUser.UnBoundIngotQuantity = &currentCount
				newUser.OnlineIngotTotal = &currentCount
			}
			gameUser = &newUser
			installUsers = append(installUsers, newUser)
			continue
		}

		// 跳过RoleGameId为nil的情况
		if gameUser.RoleGameId == nil {
			continue
		}

		switch Action_type {
		case ActionTypeScheduled:
			// 定时任务更新
			switch update_type {
			case UpdateTypeTalisman:
				gameUser.UnBoundTalisman = &currentCount
			case UpdateTypeIngot:
				gameUser.UnBoundIngotQuantity = &currentCount
			}
			gameUser.LastSyncQueryTime = &now
			gameUser.LastSyncUpdateTime = &now

		case ActionTypeManual:
			// 用户手动更新
			switch update_type {
			case UpdateTypeTalisman:
				gameUser.OnlineTalismanTotal = &currentCount
			case UpdateTypeIngot:
				gameUser.OnlineIngotTotal = &currentCount
			}
			gameUser.LastOnlineQueryTime = &now
		}

		updatedUsers = append(updatedUsers, *gameUser)
	}

	// 在这里将进行批量更新和插入
	err := update_data(installUsers, updatedUsers)
	if err != nil {
		return err
	}
	return nil

}

func update_data(installUsers []sc.GameUser, updatedUsers []sc.GameUser) error {
	// 获取数据库连接
	db := global.GVA_DB
	if db == nil {
		return fmt.Errorf("数据库连接未初始化")
	}

	// 批量插入新增用户
	if len(installUsers) > 0 {
		if err := db.CreateInBatches(&installUsers, 100).Error; err != nil {
			return fmt.Errorf("批量插入失败: %v", err)
		}
	}

	// 批量更新现有用户
	// if len(updatedUsers) > 0 {
	// 	for _, user := range updatedUsers {
	// 		if err := db.Model(&sc.GameUser{}).
	// 			Where("id = ?", user.ID).
	// 			Updates(&user).Error; err != nil {
	// 			return fmt.Errorf("更新用户 %s 失败: %v", *user.RoleGameId, err)
	// 		}
	// 	}
	// }
	err := update_user_info(updatedUsers)
	if err != nil {
		return err
	}
	return nil
}

// 批量更新用户表（单条 SQL，事务安全）
// 批量更新用户表（极简、无事务、单条 SQL）
func update_user_info(users []sc.GameUser) error {
	if len(users) == 0 {
		return nil
	}

	data := make([]map[string]interface{}, len(users))
	for i, u := range users {
		data[i] = map[string]interface{}{
			"id":                       u.ID,
			"last_ingot_trade_time":    u.LastIngotTradeTime,
			"last_talisman_trade_time": u.LastTalismanTradeTime,
			"online_talisman_total":    u.OnlineTalismanTotal,
			"online_ingot_total":       u.OnlineIngotTotal,
			"last_sync_query_time":     u.LastSyncQueryTime,
			"last_sync_update_time":    u.LastSyncUpdateTime,
			"talisman_diff":            u.TalismanDiff,
			"ingot_diff":               u.IngotDiff,
			"un_bound_ingot_quantity":  u.UnBoundIngotQuantity,
			"bound_ingot_quantity":     u.BoundIngotQuantity,
			"total_ingot_quantity":     u.TotalIngotQuantity,
			"un_bound_talisman":        u.UnBoundTalisman,
			"bound_talisman":           u.BoundTalisman,
			"total_talisman":           u.TotalTalisman,
		}
	}
	return global.GVA_DB.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}}, // 主键/唯一键
			UpdateAll: true,                          // 冲突时全部字段更新
		},
	).CreateInBatches(users, 100).Error
}
