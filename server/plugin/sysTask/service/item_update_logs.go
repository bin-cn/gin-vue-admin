package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	sc "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	"github.com/flipped-aurora/gin-vue-admin/server/service/gameManage"
	service "github.com/flipped-aurora/gin-vue-admin/server/service/smartcreate"

	"gorm.io/gorm"
)

type ItemUpdateLogs struct{}

func verifyItemUpdateLogsData(arg interface{}) (*gameManage.GameManager, error) {

	global.GVA_LOG.Info("获取物品更新日志任务开始")

	id := arg.(string)
	var cookieData *string = nil
	// 直接创建实例（不需要New方法）
	service := service.CookieDataService{}

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

// @description: 物品更新日志
// @param: arg interface{}
// @return: error
func (t ItemUpdateLogs) Exec(arg interface{}) error {
	//arg 是字符串用逗号分隔的,需要解析为数组对象
	argStr := arg.(string)
	argArr := strings.Split(argStr, ",")
	if len(argArr) != 2 {
		return fmt.Errorf("物品更新日志任务参数错误,参数格式: 任务ID,角色ID")
	}
	cookieID := argArr[0]
	buildTileId := argArr[1]

	gm, err := verifyItemUpdateLogsData(cookieID)
	if err != nil {
		return fmt.Errorf("物品更新日志任务验证数据失败: %v", err)
	}
	// 1.查询game_user_auth_code表中的所有main_server_zone字段值
	var mainServerZones []string
	result := global.GVA_DB.Raw("SELECT main_server_zone FROM gva.game_user_auth_code group by main_server_zone order by main_server_zone asc;").Scan(&mainServerZones)
	if result.Error != nil {
		return fmt.Errorf("查询game_user_auth_code表中的所有main_server_zone字段值失败: %v", result.Error)
	}

	// 转换一下类型
	minutes, err := strconv.ParseInt(buildTileId, 10, 32)
	if err != nil {
		return fmt.Errorf("buildTileId转换失败: %v", err)
	}
	startTime, endTime := gameManage.BuildTime(int(minutes))

	var game_user_auth_codes []smartcreate.GameUserAuthCode

	//1.1 查询game_user_auth_code表中的所有数据.
	db := global.GVA_DB.Model(&smartcreate.GameUserAuthCode{})
	err = db.Find(&game_user_auth_codes).Error
	if err != nil {
		return fmt.Errorf("查询game_user_auth_code表中的所有数据失败: %v", err)
	}

	//1.2查询出u_game_user表中的所有数据.
	var gameUsers []smartcreate.GameUser

	//2. mainServerZones    这里拿到的是区服ID列表，需要从game_server表中查询对应的主区服ID,需要去重

	var mainServerZoneIds []int
	result = global.GVA_DB.Raw("SELECT id FROM gva.game_server WHERE main_server_zone IN ?", mainServerZones).Scan(&mainServerZoneIds)
	if result.Error != nil {
		return fmt.Errorf("查询game_server表中的所有main_server_zone字段值失败: %v", result.Error)
	}
	// 去重	 mainServerZoneIds
	// 手动实现去重逻辑
	uniqueIds := make([]int, 0)
	idMap := make(map[int]bool)
	for _, id := range mainServerZoneIds {
		if _, exists := idMap[id]; !exists {
			idMap[id] = true
			uniqueIds = append(uniqueIds, id)
		}
	}
	mainServerZoneIds = uniqueIds

	// 3..遍历mainServerZones，调用gm.GetPlayerActions()方法
	for _, mainServerZone := range mainServerZoneIds {

		mainServerZoneStr := strconv.Itoa(mainServerZone)

		var playerActions []gameManage.PlayerAction
		var err error
		playerActions, err = gm.GetPlayerActions(startTime, endTime, "20:元宝", mainServerZoneStr)
		if err != nil {
			return fmt.Errorf("获取玩家操作失败: %v", err)
		}

		db = global.GVA_DB.Model(&smartcreate.GameUser{})
		err = db.Find(&gameUsers).Error
		if err != nil {
			return fmt.Errorf("查询gameUsers表中的所有数据失败: %v", err)
		}

		diff_update_data(playerActions, game_user_auth_codes, gameUsers, UpdateTypeIngot, ActionTypeScheduled)

		playerActions, err = gm.GetPlayerActions(startTime, endTime, "2:灵符", mainServerZoneStr)
		if err != nil {
			return fmt.Errorf("获取玩家操作失败: %v", err)
		}

		db = global.GVA_DB.Model(&smartcreate.GameUser{})
		err = db.Find(&gameUsers).Error
		if err != nil {
			return fmt.Errorf("查询gameUsers表中的所有数据失败: %v", err)
		}

		diff_update_data(playerActions, game_user_auth_codes, gameUsers, UpdateTypeIngot, ActionTypeScheduled)

	}

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
func diff_update_data(
	playerActions []gameManage.PlayerAction,
	game_user_auth_codes []sc.GameUserAuthCode,
	gameUsers []sc.GameUser,
	update_type int,
	Action_type int,
) []sc.GameUser { // 返回更新后的gameUsers
	//前提说明 玩家角色ID:  PlayerAction PlayerId 这是玩家角色ID, 对应着 GameUser中的RoleGameId   对应着 GameUserAuthCode的RoleGameId
	//  PlayerAction 中的 ItemId 代表物品ID，20是非绑定元宝 ,2 是非绑定灵符
	// update_type  2 表示本次更新的是非绑定灵符， 20表示本次更新的是非绑定元宝

	//Action_type  1 表示为定时任务更新, 2.表示为用户手动更新 ,需要更具ItemId类型来判断来更新 CurrentCount 是当前数据 来源于PlayerAction数据中
	//             如果是用户手动更新,则要更新这几个数据  OnlineTalismanTotal : 线上灵符总数    OnlineIngotTotal :线上元宝总数     LastOnlineQueryTime:是最后线上查询时间
	//            如果是定时任务更新，则要更新这几个数据  UnBoundTalisman : 未绑定灵符总数    UnBoundIngotQuantity :未绑定元宝总数  TalismanDiff：CurrentCount对比UnBoundTalisman 增加或减少的数量  IngotDiff 未CurrentCount对比UnBoundIngotQuantity增加或减少的数量
	//                                                 LastSyncQueryTime :最后同步查询时间  LastSyncUpdateTime :最后同步更新时间

	// 本次是为了更新非绑定元宝 和 非绑定灵符信息,最终是为了更新 GameUser 表中的数据

	// 1.playerActions以这个数组的角色ID为准，首先匹配 GameUser中 中的数据，如果有该角色ID存在则进行更新，
	// 				如果不存在,则需要从GameUserAuthCode表中查询数据过来,如果GameUserAuthCode表中也没有,则需要跳过，如果有,则需要插入gameUsers中的数据，本身gameUsers中的数据也是从GameUserAuthCode表中查询过来的

	//2.需要注意的是 PlayerAction中的PlayerId，也就是角色ID,可能会查出来多条角色ID一样的数据,因为这本身就是更新记录,如果有多条,则以Time时间最新的为准

	// 1. 数据流向确认，一次性即可，如果GameUser表中已经有这个角色ID了，不需要再从GameUserAuthCode更新他的信息。
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
			ServerZoneId := string(*authCode.GameServerId)
			newUser := sc.GameUser{
				RoleGameId:           authCode.RoleGameId,
				NickName:             authCode.AssignerName, // 使用游戏角色名称作为昵称
				LoginCode:            authCode.LoginCode,
				RoleGameName:         authCode.RoleGameName,
				GameServerName:       authCode.GameServerName,
				GameServerId:         authCode.GameServerId,
				ServerZoneId:         &ServerZoneId, //
				UnBoundIngotQuantity: new(int),      // 默认值0
				BoundIngotQuantity:   new(int),      // 默认值0
				TotalIngotQuantity:   new(int),      // 默认值0
				UnBoundTalisman:      new(int),      // 默认值0
				BoundTalisman:        new(int),      // 默认值0
				TotalTalisman:        new(int),      // 默认值0
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
			if update_type == UpdateTypeTalisman {
				gameUser.UnBoundTalisman = &currentCount
			} else if update_type == UpdateTypeIngot {
				gameUser.UnBoundIngotQuantity = &currentCount
			}
			gameUser.LastSyncQueryTime = &now
			gameUser.LastSyncUpdateTime = &now

		case ActionTypeManual:
			// 用户手动更新
			if update_type == UpdateTypeTalisman {
				gameUser.OnlineTalismanTotal = &currentCount
			} else if update_type == UpdateTypeIngot {
				gameUser.OnlineIngotTotal = &currentCount
			}
			gameUser.LastOnlineQueryTime = &now
		}

		updatedUsers = append(updatedUsers, *gameUser)
	}

	// 在这里将进行批量更新和插入
	update_data(installUsers, updatedUsers)
	return updatedUsers
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
	if len(updatedUsers) > 0 {
		for _, user := range updatedUsers {
			if err := db.Model(&sc.GameUser{}).
				Where("id = ?", user.ID).
				Updates(&user).Error; err != nil {
				return fmt.Errorf("更新用户 %s 失败: %v", *user.RoleGameId, err)
			}
		}
	}
	return nil
}
