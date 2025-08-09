package smartcreate

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
)

type GameServerService struct{}

// CreateGameServer 创建游戏服务器列表记录
// Author [yourname](https://github.com/yourname)
func (gamesrvService *GameServerService) CreateGameServer(ctx context.Context, gamesrv *smartcreate.GameServer) (err error) {
	err = global.GVA_DB.Create(gamesrv).Error
	return err
}

// DeleteGameServer 删除游戏服务器列表记录
// Author [yourname](https://github.com/yourname)
func (gamesrvService *GameServerService) DeleteGameServer(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&smartcreate.GameServer{}, "id = ?", ID).Error
	return err
}

// DeleteGameServerByIds 批量删除游戏服务器列表记录
// Author [yourname](https://github.com/yourname)
func (gamesrvService *GameServerService) DeleteGameServerByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]smartcreate.GameServer{}, "id in ?", IDs).Error
	return err
}

// UpdateGameServer 更新游戏服务器列表记录
// Author [yourname](https://github.com/yourname)
func (gamesrvService *GameServerService) UpdateGameServer(ctx context.Context, gamesrv smartcreate.GameServer) (err error) {
	err = global.GVA_DB.Model(&smartcreate.GameServer{}).Where("id = ?", gamesrv.ID).Updates(&gamesrv).Error
	return err
}

// GetGameServer 根据ID获取游戏服务器列表记录
// Author [yourname](https://github.com/yourname)
func (gamesrvService *GameServerService) GetGameServer(ctx context.Context, ID string) (gamesrv smartcreate.GameServer, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&gamesrv).Error
	return
}

func (gamesrvService *GameServerService) SelectAll(ctx context.Context) (list []smartcreate.GameServer, err error) {
	err = global.GVA_DB.Find(&list).Error
	return list, err
}

// GetGameServerInfoList 分页获取游戏服务器列表记录
// Author [yourname](https://github.com/yourname)
func (gamesrvService *GameServerService) GetGameServerInfoList(ctx context.Context, info smartcreateReq.GameServerSearch) (list []smartcreate.GameServer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&smartcreate.GameServer{})
	var gamesrvs []smartcreate.GameServer
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.MainServerZoneId != nil && *info.MainServerZoneId != "" {
		db = db.Where("main_server_zone_id LIKE ?", "%"+*info.MainServerZoneId+"%")
	}
	if info.MainServerZoneId != nil && *info.MainServerZoneId != "" {
		db = db.Where("main_server_zone_id LIKE ?", "%"+*info.MainServerZoneId+"%")
	}
	if info.ServerId != nil && *info.ServerId != "" {
		db = db.Where("server_id LIKE ?", "%"+*info.ServerId+"%")
	}
	if info.ServerZoneId != nil && *info.ServerZoneId != "" {
		db = db.Where("server_zone_id LIKE ?", "%"+*info.ServerZoneId+"%")
	}
	if info.ServerName != nil && *info.ServerName != "" {
		db = db.Where("server_name LIKE ?", "%"+*info.ServerName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["main_server_zone_id"] = true
	orderMap["main_server_zone_id"] = true
	orderMap["server_id"] = true
	orderMap["server_zone_id"] = true
	orderMap["server_name"] = true
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

	err = db.Find(&gamesrvs).Error
	return gamesrvs, total, err
}
func (gamesrvService *GameServerService) GetGameServerPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
