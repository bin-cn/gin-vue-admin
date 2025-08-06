package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	globalTask "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/global"
	task "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model"
	taskReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type SysTaskApi struct {
}

// CreateSysTask 创建sysTask表
// @Tags SysTask
// @Summary 创建sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.SysTask true "创建sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysTask/createSysTask [post]
func (sysTaskApi *SysTaskApi) CreateSysTask(c *gin.Context) {
	var sysTask task.SysTask
	err := c.ShouldBindJSON(&sysTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.SysTaskServiceApp.CreateSysTask(c, &sysTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysTask 删除sysTask表
// @Tags SysTask
// @Summary 删除sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.SysTask true "删除sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysTask/deleteSysTask [delete]
func (sysTaskApi *SysTaskApi) DeleteSysTask(c *gin.Context) {
	ID := c.Query("ID")
	if err := service.SysTaskServiceApp.DeleteSysTask(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysTaskByIds 批量删除sysTask表
// @Tags SysTask
// @Summary 批量删除sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysTask/deleteSysTaskByIds [delete]
func (sysTaskApi *SysTaskApi) DeleteSysTaskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := service.SysTaskServiceApp.DeleteSysTaskByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysTask 更新sysTask表
// @Tags SysTask
// @Summary 更新sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.SysTask true "更新sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysTask/updateSysTask [put]
func (sysTaskApi *SysTaskApi) UpdateSysTask(c *gin.Context) {
	var sysTask task.SysTask
	err := c.ShouldBindJSON(&sysTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.SysTaskServiceApp.UpdateSysTask(sysTask); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		if sysTask.Status == globalTask.Enabled {
			service.SysTaskServiceApp.Remove(strconv.Itoa(sysTask.ID))
			service.SysTaskServiceApp.AddTaskToTimer(strconv.Itoa(sysTask.ID))
		} else {
			service.SysTaskServiceApp.Remove(strconv.Itoa(sysTask.ID))
		}
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysTask 用id查询sysTask表
// @Tags SysTask
// @Summary 用id查询sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query task.SysTask true "用id查询sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysTask/findSysTask [get]
func (sysTaskApi *SysTaskApi) FindSysTask(c *gin.Context) {
	ID := c.Query("ID")
	if resysTask, err := service.SysTaskServiceApp.GetSysTask(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysTask": resysTask}, c)
	}
}

// GetSysTaskList 分页获取sysTask表列表
// @Tags SysTask
// @Summary 分页获取sysTask表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query taskReq.SysTaskSearch true "分页获取sysTask表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysTask/getSysTaskList [get]
func (sysTaskApi *SysTaskApi) GetSysTaskList(c *gin.Context) {
	var pageInfo taskReq.SysTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := service.SysTaskServiceApp.GetSysTaskInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// RunOne 立即执行任务
func (sysTaskApi *SysTaskApi) RunOneSysTask(c *gin.Context) {
	id := c.Query("ID")
	detail, err := service.SysTaskServiceApp.GetSysTask(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	detail.CronExpression = "手动运行"
	service.SysTaskServiceApp.Run(detail)
	response.OkWithMessage("任务已开始运行,请查看结果", c)
}

// EnableSysTask 激活任务
func (sysTaskApi *SysTaskApi) EnableSysTask(c *gin.Context) {
	id := c.Query("ID")
	status := service.SysTaskServiceApp.ChangeStatus(id, globalTask.Enabled)
	response.OkWithData(status, c)
}

// DisableSysTask 暂停任务
func (sysTaskApi *SysTaskApi) DisableSysTask(c *gin.Context) {
	id := c.Query("ID")
	status := service.SysTaskServiceApp.ChangeStatus(id, globalTask.Disabled)
	response.OkWithData(status, c)
}

// RemoveSysTask 删除任务
func (sysTaskApi *SysTaskApi) RemoveSysTask(c *gin.Context) {
	id := c.Query("ID")
	err := service.SysTaskServiceApp.DeleteSysTask(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service.SysTaskServiceApp.Remove(id)
	response.OkWithMessage("任务删除成功", c)
}
