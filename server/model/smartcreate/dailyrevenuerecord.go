// 自动生成模板DailyRevenueRecord
package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 日收入统计 结构体  DailyRevenueRecord
type DailyRevenueRecord struct {
	global.GVA_MODEL
	StatisticDate    *string `json:"statisticDate" form:"statisticDate" gorm:"column:statistic_date;" binding:"required"`            //统计日期
	UserId           *string `json:"userId" form:"userId" gorm:"column:user_id;" binding:"required"`                                 //用户ID
	UserNickname     *string `json:"userNickname" form:"userNickname" gorm:"column:user_nickname;" binding:"required"`               //用户昵称
	ServerName       *string `json:"serverName" form:"serverName" gorm:"column:server_name;" binding:"required"`                     //区服名字
	ServerZoneId     *string `json:"serverZoneId" form:"serverZoneId" gorm:"column:server_zone_id;" binding:"required"`              //区服ID
	MainServerZoneId *string `json:"mainServerZoneId" form:"mainServerZoneId" gorm:"column:main_server_zone_id;" binding:"required"` //主区服ID
	StatisticType    *string `json:"statisticType" form:"statisticType" gorm:"column:statistic_type;" binding:"required"`            //统计类型
	Amount           *int64  `json:"amount" form:"amount" gorm:"column:amount;" binding:"required"`                                  //产出金额
}

// TableName 日收入统计 DailyRevenueRecord自定义表名 daily_revenue_records
func (DailyRevenueRecord) TableName() string {
	return "daily_revenue_records"
}
