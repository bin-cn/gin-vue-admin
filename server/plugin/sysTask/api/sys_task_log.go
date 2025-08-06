package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	task "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model"
	taskReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysTaskLogApi struct {
}

// CreateSysTaskLog 创建sysTaskLog表
// @Tags SysTaskLog
// @Summary 创建sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.SysTaskLog true "创建sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysTaskLog/createSysTaskLog [post]
func (sysTaskLogApi *SysTaskLogApi) CreateSysTaskLog(c *gin.Context) {
	var sysTaskLog task.SysTaskLog
	err := c.ShouldBindJSON(&sysTaskLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.SysTaskLogServiceApp.CreateSysTaskLog(&sysTaskLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysTaskLog 删除sysTaskLog表
// @Tags SysTaskLog
// @Summary 删除sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.SysTaskLog true "删除sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysTaskLog/deleteSysTaskLog [delete]
func (sysTaskLogApi *SysTaskLogApi) DeleteSysTaskLog(c *gin.Context) {
	ID := c.Query("ID")
	if err := service.SysTaskLogServiceApp.DeleteSysTaskLog(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysTaskLogByIds 批量删除sysTaskLog表
// @Tags SysTaskLog
// @Summary 批量删除sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysTaskLog/deleteSysTaskLogByIds [delete]
func (sysTaskLogApi *SysTaskLogApi) DeleteSysTaskLogByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := service.SysTaskLogServiceApp.DeleteSysTaskLogByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysTaskLog 更新sysTaskLog表
// @Tags SysTaskLog
// @Summary 更新sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.SysTaskLog true "更新sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysTaskLog/updateSysTaskLog [put]
func (sysTaskLogApi *SysTaskLogApi) UpdateSysTaskLog(c *gin.Context) {
	var sysTaskLog task.SysTaskLog
	err := c.ShouldBindJSON(&sysTaskLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.SysTaskLogServiceApp.UpdateSysTaskLog(sysTaskLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysTaskLog 用id查询sysTaskLog表
// @Tags SysTaskLog
// @Summary 用id查询sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query task.SysTaskLog true "用id查询sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysTaskLog/findSysTaskLog [get]
func (sysTaskLogApi *SysTaskLogApi) FindSysTaskLog(c *gin.Context) {
	ID := c.Query("ID")
	if resysTaskLog, err := service.SysTaskLogServiceApp.GetSysTaskLog(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysTaskLog": resysTaskLog}, c)
	}
}

// GetSysTaskLogList 分页获取sysTaskLog表列表
// @Tags SysTaskLog
// @Summary 分页获取sysTaskLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query taskReq.SysTaskLogSearch true "分页获取sysTaskLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysTaskLog/getSysTaskLogList [get]
func (sysTaskLogApi *SysTaskLogApi) GetSysTaskLogList(c *gin.Context) {
	var pageInfo taskReq.SysTaskLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := service.SysTaskLogServiceApp.GetSysTaskLogInfoList(pageInfo); err != nil {
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
