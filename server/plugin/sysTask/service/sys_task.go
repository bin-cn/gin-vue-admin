package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	taskGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/global"
	task "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model"
	taskReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SysTaskService struct {
}

var SysTaskServiceApp = new(SysTaskService)

// CreateSysTask 创建sysTask表记录
func (sysTaskService *SysTaskService) CreateSysTask(c *gin.Context, sysTask *task.SysTask) (err error) {
	err = global.GVA_DB.Create(sysTask).Error
	if err == nil && sysTask.Status == taskGlobal.Enabled {
		SysTaskServiceApp.ChangeStatus(strconv.Itoa(sysTask.ID), taskGlobal.Enabled)
	}
	return err
}

// DeleteSysTask 删除sysTask表记录
func (sysTaskService *SysTaskService) DeleteSysTask(ID string) (err error) {
	err = global.GVA_DB.Delete(&task.SysTask{}, "id = ?", ID).Error
	return err
}

// DeleteSysTaskByIds 批量删除sysTask表记录
func (sysTaskService *SysTaskService) DeleteSysTaskByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]task.SysTask{}, "id in ?", IDs).Error
	return err
}

// UpdateSysTask 更新sysTask表记录
func (sysTaskService *SysTaskService) UpdateSysTask(sysTask task.SysTask) (err error) {

	err = global.GVA_DB.Save(&sysTask).Error
	return err
}

// GetSysTask 根据ID获取sysTask表记录
func (sysTaskService *SysTaskService) GetSysTask(ID string) (sysTask task.SysTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysTask).Error
	return
}

// GetSysTaskInfoList 分页获取sysTask表记录
func (sysTaskService *SysTaskService) GetSysTaskInfoList(info taskReq.SysTaskSearch) (list []task.SysTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&task.SysTask{})
	var sysTasks []task.SysTask
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.InvokeTarget != "" {
		db = db.Where("invoke_target like ?", "%"+info.InvokeTarget+"%")
	}
	if info.Cron != "" {
		db = db.Where("cron_expression = ?", info.Cron)
	}
	if info.Group != "" {
		db = db.Where("task_group = ?", info.Group)
	}
	if info.Remark != "" {
		db = db.Where("remark like ?", "%"+info.Remark+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.Tag != "" {
		db = db.Where("tag like ? ", "%"+info.Tag+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysTasks).Error
	return sysTasks, total, err
}

// UpdateTaskStatus 根据id修改任务状态
func (sysTaskService *SysTaskService) UpdateTaskStatus(id int, status taskGlobal.SysTaskStatus) (err error) {
	taskForm := new(task.SysTask)
	if err = global.GVA_DB.Where("id=?", id).First(taskForm).Error; err != nil {
		return err
	}
	taskForm.Status = status
	return global.GVA_DB.Save(taskForm).Error
}

// ActiveSysTaskList 获取所有激活任务
func (sysTaskService *SysTaskService) ActiveSysTaskList() ([]task.SysTask, error) {
	var tasks []task.SysTask
	err := global.GVA_DB.Where("status = ?", taskGlobal.Enabled).Find(&tasks).Error
	return tasks, err
}
