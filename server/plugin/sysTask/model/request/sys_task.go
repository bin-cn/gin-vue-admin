package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SysTaskSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	InvokeTarget   string     `json:"invokeTarget" form:"invokeTarget" `
	Group          string     `json:"group" form:"group"`
	Cron           string     `json:"cron" form:"cron"`
	Remark         string     `json:"remark" form:"remark"`
	Status         string     `json:"status" form:"status"`
	Tag            string     `json:"tag" form:"tag"` //tag

	request.PageInfo
}
