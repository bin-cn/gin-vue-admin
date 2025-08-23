package smartcreate

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/response"
)

type DailyRevenueRecordService struct{}

// CreateDailyRevenueRecord 创建日收入统计记录
// Author [yourname](https://github.com/yourname)
func (drService *DailyRevenueRecordService) CreateDailyRevenueRecord(ctx context.Context, dr *smartcreate.DailyRevenueRecord) (err error) {
	err = global.GVA_DB.Create(dr).Error
	return err
}

func (drService *DailyRevenueRecordService) InstallDailyRevenueRecord() (err error) {
	yesterdayStr := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	// 1. 昨天是否已有数据
	var cnt int64
	if err = global.GVA_DB.Model(&smartcreate.DailyRevenueRecord{}).
		Where("statistic_date = ?", yesterdayStr).
		Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return nil // 已有数据，跳过
	}

	// 2. 统计昨日数据
	type Agg struct {
		GameServerName    *string
		ServerZoneID      *string
		NickName          *string
		GameServerID      *string
		UserId            *string
		TotalUnBoundIngot *int64
	}
	var list []Agg

	sql := `
		SELECT  
		MAX(t1.game_server_name) AS game_server_name,  
		MAX(t1.server_zone_id) AS server_zone_id,    
		max(t1.nick_name) AS nick_name,
		t1.game_server_id game_server_id,
		max(t1.user_id ) as user_id,
		SUM(t1.un_bound_ingot_quantity) AS total_un_bound_ingot 
		FROM u_game_user AS t1
		GROUP BY t1.game_server_id 
		ORDER BY game_server_name ASC;`

	if err = global.GVA_DB.Raw(sql).Scan(&list).Error; err != nil {
		return err
	}

	statisticDate := &yesterdayStr

	// 3. 批量插入
	records := make([]smartcreate.DailyRevenueRecord, 0, len(list))
	for _, v := range list {
		records = append(records, smartcreate.DailyRevenueRecord{
			StatisticDate:    statisticDate,
			ServerName:       v.GameServerName,
			UserId:           v.UserId,
			UserNickname:     v.NickName,
			ServerZoneId:     v.GameServerID,
			MainServerZoneId: v.ServerZoneID,
			Amount:           v.TotalUnBoundIngot,
		})
	}

	return global.GVA_DB.Create(&records).Error
}

// DeleteDailyRevenueRecord 删除日收入统计记录
// Author [yourname](https://github.com/yourname)
func (drService *DailyRevenueRecordService) DeleteDailyRevenueRecord(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&smartcreate.DailyRevenueRecord{}, "id = ?", ID).Error
	return err
}

// DeleteDailyRevenueRecordByIds 批量删除日收入统计记录
// Author [yourname](https://github.com/yourname)
func (drService *DailyRevenueRecordService) DeleteDailyRevenueRecordByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]smartcreate.DailyRevenueRecord{}, "id in ?", IDs).Error
	return err
}

// UpdateDailyRevenueRecord 更新日收入统计记录
// Author [yourname](https://github.com/yourname)
func (drService *DailyRevenueRecordService) UpdateDailyRevenueRecord(ctx context.Context, dr smartcreate.DailyRevenueRecord) (err error) {
	err = global.GVA_DB.Model(&smartcreate.DailyRevenueRecord{}).Where("id = ?", dr.ID).Updates(&dr).Error
	return err
}

// GetDailyRevenueRecord 根据ID获取日收入统计记录
// Author [yourname](https://github.com/yourname)
func (drService *DailyRevenueRecordService) GetDailyRevenueRecord(ctx context.Context, ID string) (dr smartcreate.DailyRevenueRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dr).Error
	return
}

// GetDailyRevenueRecordInfoList 分页获取日收入统计记录
// Author [yourname](https://github.com/yourname)
func (drService *DailyRevenueRecordService) GetDailyRevenueRecordInfoList(ctx context.Context, info smartcreateReq.DailyRevenueRecordSearch) (list []smartcreate.DailyRevenueRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&smartcreate.DailyRevenueRecord{})
	var drs []smartcreate.DailyRevenueRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["statistic_date"] = true
	orderMap["user_id"] = true
	orderMap["user_nickname"] = true
	orderMap["server_name"] = true
	orderMap["server_zone_id"] = true
	orderMap["main_server_zone_id"] = true
	orderMap["statistic_type"] = true
	orderMap["amount"] = true
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

	err = db.Find(&drs).Error
	return drs, total, err
}

// 统计数据
func (drService *DailyRevenueRecordService) Statistic(ctx context.Context, info smartcreateReq.DailyRevenueStatisticSearch) (list []response.DailyRevenueStatistic, err error) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
	sql_str := ""
	switch info.StatisticType {
	case 1:
		// 用户统计
		user_statistic_sql := `SELECT  
				MAX(user_nickname) AS user_nickname,
				user_id ,           
				MAX(server_name) AS server_name,
				SUBSTRING(statistic_date, 9, 2) AS statistic_date,
				MAX(server_zone_id) AS server_zone_id,  
				SUM(amount) AS amount            
				FROM gva.daily_revenue_records 
				WHERE statistic_date LIKE CONCAT(?, '%')
				GROUP BY user_id, statistic_date
				ORDER BY statistic_date, user_id`
		sql_str = user_statistic_sql
	case 2:
		// 区服统计
		server_statistic_sql := `SELECT
				MAX(user_nickname) AS user_nickname,  
				MAX(user_id) AS user_id,              
				MAX(server_name) AS server_name,
				SUBSTRING(statistic_date, 9, 2) AS statistic_date,
				server_zone_id,                       
				SUM(amount) AS amount           
				FROM gva.daily_revenue_records 
				WHERE statistic_date LIKE CONCAT(?, '%')
				GROUP BY server_zone_id, statistic_date
				ORDER BY statistic_date, server_zone_id`
		sql_str = server_statistic_sql
	}
	var statisticList []response.DailyRevenueStatistic
	err = global.GVA_DB.Raw(sql_str, info.StatisticDate).Scan(&statisticList).Error
	return statisticList, err
}

func (drService *DailyRevenueRecordService) GetDailyRevenueRecordPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
