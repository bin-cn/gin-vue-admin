package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"

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

// ----------------------------->更新服务器ID<-----------------------------
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

	// 按区服ID分组服务器
	zoneServers := make(map[string][]model.GameServer)
	for _, server := range list {
		if server.ServerZoneId != nil {
			zone := *server.ServerZoneId
			zoneServers[zone] = append(zoneServers[zone], server)
		}
	}

	// 收集需要更新的服务器
	var serversToUpdate []model.GameServer
	updateCount := 0

	// 从serverName中提取_nX后缀数字的函数
	extractNumber := func(serverName string) int {
		// 查找最后一个下划线和n的组合
		re := regexp.MustCompile(`_n(\d+)$`)
		matches := re.FindStringSubmatch(serverName)
		if len(matches) > 1 {
			if num, err := strconv.Atoi(matches[1]); err == nil {
				return num
			}
		}
		return 0 // 默认为0
	}

	// 处理合区信息
	for _, mergeInfo := range margeInfo {
		// 解析被合区的区服ID列表
		mergedZones := strings.Split(mergeInfo.MergedZones, ",")
		mainZone := mergeInfo.Zone

		// 处理主区服
		if servers, exists := zoneServers[mainZone]; exists {
			// 在主区服内选择主服务器（根据_nX后缀数字最大的）
			var mainServer model.GameServer
			maxNumber := -1

			for _, server := range servers {
				if server.ServerName != nil {
					num := extractNumber(*server.ServerName)
					if num > maxNumber {
						maxNumber = num
						mainServer = server
					}
				}
			}

			if maxNumber >= 0 {
				// 设置主服务器
				newMainZoneId := *mainServer.ServerZoneId
				newMainServerId := ""
				if mainServer.ServerId != nil {
					newMainServerId = *mainServer.ServerId
				}

				// 更新主服务器自身
				if mainServer.MainServerZoneId == nil || *mainServer.MainServerZoneId != newMainZoneId ||
					mainServer.MainServerId == nil || *mainServer.MainServerId != newMainServerId {
					updatedServer := mainServer
					updatedServer.MainServerZoneId = &newMainZoneId
					updatedServer.MainServerId = &newMainServerId
					serversToUpdate = append(serversToUpdate, updatedServer)
					updateCount++
				}

				// 更新同一区服的其他服务器
				for _, server := range servers {
					if server.ServerId != nil && (mainServer.ServerId == nil || *server.ServerId != *mainServer.ServerId) {
						if server.MainServerZoneId == nil || *server.MainServerZoneId != newMainZoneId ||
							server.MainServerId == nil || *server.MainServerId != newMainServerId {
							updatedServer := server
							updatedServer.MainServerZoneId = &newMainZoneId
							updatedServer.MainServerId = &newMainServerId
							serversToUpdate = append(serversToUpdate, updatedServer)
							updateCount++
						}
					}
				}
			}
		}

		// 处理被合区的服务器
		for _, zoneStr := range mergedZones {
			zoneStr = strings.TrimSpace(zoneStr)
			if zoneStr == "" || zoneStr == mainZone {
				continue
			}

			if servers, exists := zoneServers[zoneStr]; exists {
				// 被合区的服务器主服务器ID设置为主区服的主服务器ID
				for _, server := range servers {
					// 找到主区服的主服务器
					var mainServer model.GameServer
					maxNumber := -1
					if mainServers, mainExists := zoneServers[mainZone]; mainExists {
						for _, ms := range mainServers {
							if ms.ServerName != nil {
								num := extractNumber(*ms.ServerName)
								if num > maxNumber {
									maxNumber = num
									mainServer = ms
								}
							}
						}
					}

					if maxNumber >= 0 {
						newMainZoneId := mainZone
						newMainServerId := ""
						if mainServer.ServerId != nil {
							newMainServerId = *mainServer.ServerId
						}

						if server.MainServerZoneId == nil || *server.MainServerZoneId != newMainZoneId ||
							server.MainServerId == nil || *server.MainServerId != newMainServerId {
							updatedServer := server
							updatedServer.MainServerZoneId = &newMainZoneId
							updatedServer.MainServerId = &newMainServerId
							serversToUpdate = append(serversToUpdate, updatedServer)
							updateCount++
						}
					}
				}
			}
		}
	}

	// 处理没有合区信息的独立区服
	for zoneId, servers := range zoneServers {
		// 检查该区服是否在合区信息中
		inMergeInfo := false
		for _, mergeInfo := range margeInfo {
			if mergeInfo.Zone == zoneId {
				inMergeInfo = true
				break
			}
			mergedZones := strings.Split(mergeInfo.MergedZones, ",")
			for _, zoneStr := range mergedZones {
				if strings.TrimSpace(zoneStr) == zoneId {
					inMergeInfo = true
					break
				}
			}
		}

		if !inMergeInfo && len(servers) > 1 {
			// 独立区服但有多台服务器，选择主服务器
			var mainServer model.GameServer
			maxNumber := -1

			for _, server := range servers {
				if server.ServerName != nil {
					num := extractNumber(*server.ServerName)
					if num > maxNumber {
						maxNumber = num
						mainServer = server
					}
				}
			}

			if maxNumber >= 0 {
				newMainZoneId := *mainServer.ServerZoneId
				newMainServerId := ""
				if mainServer.ServerId != nil {
					newMainServerId = *mainServer.ServerId
				}

				// 更新所有服务器
				for _, server := range servers {
					if server.MainServerZoneId == nil || *server.MainServerZoneId != newMainZoneId ||
						server.MainServerId == nil || *server.MainServerId != newMainServerId {
						updatedServer := server
						updatedServer.MainServerZoneId = &newMainZoneId
						updatedServer.MainServerId = &newMainServerId
						serversToUpdate = append(serversToUpdate, updatedServer)
						updateCount++
					}
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
				Where("server_zone_id = ? AND server_id = ?", *server.ServerZoneId, *server.ServerId).
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

			existingServers[*server.ServerId] = server
		}
	}

	// 收集需要创建的新服务器
	var newServers []model.GameServer

	// 遍历网络获取的服务器数据
	for _, newServer := range servers {
		// 如果ServerId在数据库中不存在，则准备创建新记录
		if _, exists := existingServers[newServer.Id]; !exists {
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

//----------------------------->查物品记录<-----------------------------
