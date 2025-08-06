// 自动生成模板SysTaskLog
package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/global"
	"time"
)

// sysTaskLog表 结构体  SysTaskLog
type SysTaskLog struct {
	ID             int                  `gorm:"primarykey" json:"ID"`                                                                                // 主键ID
	TaskId         int                  `json:"taskId" form:"taskId" gorm:"column:task_id;comment:任务id;size:10;"`                                    //任务id
	TaskName       string               `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:32;"`                              //任务名称
	CronExpression string               `json:"cronExpression" form:"cronExpression" gorm:"column:cron_expression;comment:cron_expression;size:64;"` //cron_expression
	InvokeTarget   string               `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:invoke_target;size:256;"`        //invoke_target
	Timeout        int                  `json:"timeout" form:"timeout" gorm:"column:timeout;comment:;"`                                              //timeout
	RetryTimes     int                  `json:"retryTimes" form:"retryTimes" gorm:"column:retry_times;comment:;"`                                    //retryTimes
	StartTime      time.Time            `json:"startTime" form:"startTime" gorm:"column:start_time;comment:;"`                                       //startTime
	EndTime        time.Time            `json:"endTime" form:"endTime" gorm:"column:end_time;comment:;"`                                             //endTime
	Status         global.SysTaskStatus `json:"status" form:"status" gorm:"column:status;comment:;"`                                                 //status
	Result         string               `json:"result" form:"result" gorm:"column:result;comment:;"`                                                 //result
	TotalTime      int                  `json:"totalTime" form:"totalTime" gorm:"column:total_time;comment:;size:19;"`                               //totalTime
}

// TableName sysTaskLog表 SysTaskLog自定义表名 sys_task_log
func (SysTaskLog) TableName() string {
	return "sys_task_log"
}
