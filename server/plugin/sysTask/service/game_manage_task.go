package service

import (
	"context"
	"errors"
	"fmt"

	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	"github.com/flipped-aurora/gin-vue-admin/server/service/gameManage"
	"github.com/flipped-aurora/gin-vue-admin/server/service/smartcreate"
	"gorm.io/gorm"
)

// MyCustomTask 获取游戏服务器列表和服务器合区信息的任务

type ListServerTask struct{}

// verifyData 前置验证数据.验证数据

func verifyData(arg interface{}) ([]gameManage.NewServer, *gameManage.GameManager, error) {
	global.GVA_LOG.Info("获取服务器列表信息和合区信息任务开始")
	id := arg.(string)
	var cookieData *string = nil
	// 直接创建实例（不需要New方法）
	service := smartcreate.CookieDataService{}

	data, err := service.GetCookieData(nil, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没查到数据
			return nil, nil, fmt.Errorf("未找到记录对应的cookie记录数据")
		}
		// 其他数据库错误
		return nil, nil, fmt.Errorf("查询错误: %v", err)
	}

	if *data.Status == 0 {
		return nil, nil, fmt.Errorf("cookie记录数据未启用")
	}

	cookieData = data.CookieValue
	gm := gameManage.NewGameManager(*cookieData)

	servers, err := gm.Operservers()
	if err != nil {
		return nil, nil, fmt.Errorf("获取服务器列表失败: %v", err)
	}

	if len(servers) == 0 {
		data.Status = new(int) // 默认为0
		service.UpdateCookieData(nil, data)
		return nil, nil, fmt.Errorf("获取服务器列表为空")
	}

	// 更新一下最后使用时间.
	now := time.Now()
	data.LastUseTime = &now
	service.UpdateCookieData(nil, data)

	return servers, gm, nil
}

// 更新主服务器ID
// 更新主服务器ID
func updateMainId(gm *gameManage.GameManager) error {
	gameService := smartcreate.GameServerService{}

	margeInfo, err := gm.GetServerList()
	if err != nil {
		return fmt.Errorf("获取合并服务器列表失败: %v", err)
	}

	list, err := gameService.SelectAll(context.Background())
	if err != nil {
		return fmt.Errorf("获取数据库服务器列表失败: %v", err)
	}

	// 创建区服ID到服务器的映射，方便快速查找
	serverMap := make(map[string]model.GameServer)
	for _, server := range list {
		if server.ServerZoneId != nil {
			serverMap[*server.ServerZoneId] = server
		}
	}

	// 收集需要更新的服务器
	var serversToUpdate []model.GameServer
	updateCount := 0

	// 遍历合区信息
	for _, mergeInfo := range margeInfo {
		// 解析被合区的区服ID列表
		mergedZones := strings.Split(mergeInfo.MergedZones, ",")

		// 获取主区服的服务器信息
		mainZone := mergeInfo.Zone
		mainServer, mainExists := serverMap[mainZone]
		if !mainExists {
			continue // 主区服不存在，跳过
		}

		// 主区服服务器：设置为自己的值
		if mainServer.ServerZoneId != nil && mainServer.ServerId != nil {
			newMainZoneId := *mainServer.ServerZoneId
			newMainServerId := *mainServer.ServerId

			// 检查是否需要更新
			currentMainZone := ""
			if mainServer.MainServerZoneId != nil {
				currentMainZone = *mainServer.MainServerZoneId
			}

			if currentMainZone != newMainZoneId {
				updatedServer := mainServer
				updatedServer.MainServerZoneId = &newMainZoneId
				updatedServer.MainServerId = &newMainServerId
				serversToUpdate = append(serversToUpdate, updatedServer)
				updateCount++
			}
		}

		// 处理被合区的服务器
		for _, zoneStr := range mergedZones {
			zoneStr = strings.TrimSpace(zoneStr)
			if zoneStr == "" {
				continue
			}

			// 跳过主区服本身
			if zoneStr == mainZone {
				continue
			}

			mergedServer, exists := serverMap[zoneStr]
			if !exists {
				continue // 被合区服务器不存在，跳过
			}

			if mergedServer.ServerZoneId != nil {
				// 被合区服务器：主区服ID设置为主区服的Zone，主服务器ID设置为主区服的主服务器ID
				newMainZoneId := mainZone
				newMainServerId := ""
				if mainServer.ServerId != nil {
					newMainServerId = *mainServer.ServerId
				}

				// 检查是否需要更新
				currentMainZone := ""
				if mergedServer.MainServerZoneId != nil {
					currentMainZone = *mergedServer.MainServerZoneId
				}

				if currentMainZone != newMainZoneId {
					updatedServer := mergedServer
					updatedServer.MainServerZoneId = &newMainZoneId
					updatedServer.MainServerId = &newMainServerId
					serversToUpdate = append(serversToUpdate, updatedServer)
					updateCount++
				}
			}
		}
	}

	// 批量更新数据库
	if len(serversToUpdate) > 0 {
		tx := global.GVA_DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		for _, server := range serversToUpdate {
			if err := tx.Model(&model.GameServer{}).
				Where("server_zone_id = ?", *server.ServerZoneId).
				Updates(map[string]interface{}{
					"main_server_zone_id": *server.MainServerZoneId,
					"main_server_id":      *server.MainServerId,
				}).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("批量更新主服务器信息失败: %v", err)
			}
		}

		tx.Commit()
		fmt.Printf("成功更新 %d 条服务器的主区服和主服务器信息", updateCount)
	}

	return nil
}

func (t ListServerTask) Exec(arg interface{}) error {
	servers, gm, err := verifyData(arg)
	if err != nil {
		return fmt.Errorf("获取服务器列表失败: %v", err)
	}

	gameService := smartcreate.GameServerService{}
	list, err := gameService.SelectAll(context.Background())
	if err != nil {
		return fmt.Errorf("获取数据库服务器列表失败: %v", err)
	}

	// 创建ServerId到GameServer的映射，便于快速查找
	existingServers := make(map[string]model.GameServer)
	for _, server := range list {
		if server.ServerZoneId != nil {
			existingServers[*server.ServerZoneId] = server
		}
	}

	// 收集需要创建的新服务器
	var newServers []model.GameServer

	// 遍历网络获取的服务器数据
	for _, newServer := range servers {
		// 如果ServerId在数据库中不存在，则准备创建新记录
		if _, exists := existingServers[newServer.ServerId]; !exists {
			gameServer := model.GameServer{
				ServerId:     &newServer.Id,
				ServerName:   &newServer.Name,
				ServerZoneId: &newServer.ServerId,
				// 根据需要设置其他字段的默认值
			}
			newServers = append(newServers, gameServer)
		}
	}

	// 批量创建新记录
	if len(newServers) > 0 {
		if err := global.GVA_DB.Create(&newServers).Error; err != nil {
			return fmt.Errorf("批量创建服务器记录失败: %v", err)
		}
		fmt.Printf("成功批量创建 %d 条新服务器记录", len(newServers))
	}

	updateMainId(gm)

	return nil
}
