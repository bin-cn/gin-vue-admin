// 自动生成模板SysTask
package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/global"
	taskGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/global"
	"gorm.io/gorm"
	"time"
)

// sysTask表 结构体  SysTask
type SysTask struct {
	ID             int                       `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt      time.Time                 // 创建时间
	UpdatedAt      time.Time                 // 更新时间
	DeletedAt      gorm.DeletedAt            `gorm:"index" json:"-"`                                                                                      // 删除时间
	TaskName       string                    `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:32;"`                              //任务名称
	TaskGroup      string                    `json:"taskGroup" form:"taskGroup" gorm:"column:task_group;comment:任务分组;size:255;"`                          //任务分组
	CronExpression string                    `json:"cronExpression" form:"cronExpression" gorm:"column:cron_expression;comment:cron_expression;size:64;"` //cron_expression
	InvokeTarget   string                    `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:;size:255;"`                     //invokeTarget
	Args           string                    `json:"args" form:"args" gorm:"column:args;comment:;size:255;"`                                              //args
	HttpMethod     taskGlobal.TaskHTTPMethod `json:"httpMethod" form:"httpMethod" gorm:"column:http_method;size:50;"`
	Timeout        int                       `json:"timeout" form:"timeout" gorm:"column:timeout;comment:任务执行超时时间(单位秒),0不限制;"`        //任务执行超时时间(单位秒),0不限制
	Multi          int                       `json:"multi" form:"multi" gorm:"column:multi;comment:是否允许多实例运行;"`                       //是否允许多实例运行
	RetryTimes     int                       `json:"retryTimes" form:"retryTimes" gorm:"column:retry_times;comment:重试次数;"`            //重试次数
	RetryInterval  int                       `json:"retryInterval" form:"retryInterval" gorm:"column:retry_interval;comment:重试间隔时间;"` //重试间隔时间
	Tag            string                    `json:"tag" form:"tag" gorm:"column:tag;comment:;size:32;"`                              //tag
	Remark         string                    `json:"remark" form:"remark" gorm:"column:remark;comment:;size:100;"`                    //remark
	Status         global.SysTaskStatus      `json:"status" form:"status" gorm:"column:status;comment:状态 1:正常 0:停止;"`                 //状态 1:正常 0:停止
}

type SysTaskResult struct {
	Result     string
	Err        error
	SysTaskLog SysTaskLog
}

// TableName sysTask表 SysTask自定义表名 sys_task
func (SysTask) TableName() string {
	return "sys_task"
}
