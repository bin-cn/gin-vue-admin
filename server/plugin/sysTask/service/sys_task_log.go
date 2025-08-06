package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	taskGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/global"
	task "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model"
	taskReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model/request"
	"time"
)

type SysTaskLogService struct {
}

var SysTaskLogServiceApp = new(SysTaskLogService)

// CreateSysTaskLog 创建sysTaskLog表记录
func (sysTaskLogService *SysTaskLogService) CreateSysTaskLog(sysTaskLog *task.SysTaskLog) (err error) {
	err = global.GVA_DB.Create(sysTaskLog).Error
	return err
}

// DeleteSysTaskLog 删除sysTaskLog表记录
func (sysTaskLogService *SysTaskLogService) DeleteSysTaskLog(ID string) (err error) {
	err = global.GVA_DB.Delete(&task.SysTaskLog{}, "id = ?", ID).Error
	return err
}

// DeleteSysTaskLogByIds 批量删除sysTaskLog表记录
func (sysTaskLogService *SysTaskLogService) DeleteSysTaskLogByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]task.SysTaskLog{}, "id in ?", IDs).Error
	return err
}

// UpdateSysTaskLog 更新sysTaskLog表记录
func (sysTaskLogService *SysTaskLogService) UpdateSysTaskLog(sysTaskLog task.SysTaskLog) (err error) {
	err = global.GVA_DB.Save(&sysTaskLog).Error
	return err
}

// GetSysTaskLog 根据ID获取sysTaskLog表记录
func (sysTaskLogService *SysTaskLogService) GetSysTaskLog(ID string) (sysTaskLog task.SysTaskLog, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysTaskLog).Error
	return
}

// GetSysTaskLogInfoList 分页获取sysTaskLog表记录
func (sysTaskLogService *SysTaskLogService) GetSysTaskLogInfoList(info taskReq.SysTaskLogSearch) (list []task.SysTaskLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&task.SysTaskLog{})
	var sysTaskLogs []task.SysTaskLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TaskName != "" {
		db = db.Where("task_name like ?", "%"+info.TaskName+"%")
	}
	if info.CronExpression != "" {
		db = db.Where("cron_expression = ?", info.CronExpression)
	}
	if info.InvokeTarget != "" {
		db = db.Where("invoke_target like ? ", "%"+info.InvokeTarget+"%")
	}
	if info.Result != "" {
		db = db.Where("result like ?", "%"+info.Result+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("start_time desc").Find(&sysTaskLogs).Error
	return sysTaskLogs, total, err
}

// CreateTaskLogByTask 根据sysTask创建任务日志
func (sysTaskLogService *SysTaskLogService) CreateTaskLogByTask(sysTask task.SysTask, status taskGlobal.SysTaskStatus) (insertId int, err error) {
	taskLogModel := new(task.SysTaskLog)
	taskLogModel.TaskId = sysTask.ID
	taskLogModel.TaskName = sysTask.TaskName
	taskLogModel.CronExpression = sysTask.CronExpression
	taskLogModel.InvokeTarget = sysTask.InvokeTarget
	taskLogModel.Timeout = sysTask.Timeout
	taskLogModel.StartTime = time.Now()
	taskLogModel.EndTime = time.Now()
	taskLogModel.Status = status
	err = global.GVA_DB.Create(&taskLogModel).Error
	if err == nil {
		insertId = taskLogModel.ID
	}
	return insertId, err
}

// UpdateTaskLogByTaskResult 更新任务日志
func (sysTaskLogService *SysTaskLogService) UpdateTaskLogByTaskResult(taskLogId int, taskResult task.SysTaskResult) (int64, error) {
	var status taskGlobal.SysTaskStatus
	result := taskResult.Result
	if taskResult.Err != nil {
		status = taskGlobal.Failure
	} else {
		status = taskGlobal.Finish
	}

	updateFields := map[string]interface{}{
		"retry_times": taskResult.SysTaskLog.RetryTimes,
		"status":      status,
		"result":      result,
		"end_time":    time.Now(),
		"total_time":  time.Now().Sub(taskResult.SysTaskLog.StartTime).Seconds(),
	}
	updates := global.GVA_DB.Model(&task.SysTaskLog{}).Where("id = ?", taskLogId).Updates(updateFields)
	if updates.Error != nil {
		return 0, updates.Error
	}
	// 返回影响的行数
	return updates.RowsAffected, nil
}
