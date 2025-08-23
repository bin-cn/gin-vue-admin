package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DailyRevenueRecordSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type DailyRevenueStatisticSearch struct {
	StatisticDate string `json:"statisticDate" form:"statisticDate" binding:"required"` //统计日期
	StatisticType int    `json:"statisticType" form:"statisticType" binding:"required"` // 统计类型
}
