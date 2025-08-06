package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SysTaskLogSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	TaskName       string     `json:"taskName" form:"taskName" `             //任务名称
	CronExpression string     `json:"cronExpression" form:"cronExpression" ` //cron_expression
	InvokeTarget   string     `json:"invokeTarget" form:"invokeTarget" `     //invoke_target
	Status         int        `json:"status" form:"status"`                  //status
	Result         string     `json:"result" form:"result" `

	request.PageInfo
}
