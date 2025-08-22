package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DailyRevenueRecordApi struct{}

// CreateDailyRevenueRecord 创建日收入统计
// @Tags DailyRevenueRecord
// @Summary 创建日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.DailyRevenueRecord true "创建日收入统计"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dr/createDailyRevenueRecord [post]
func (drApi *DailyRevenueRecordApi) CreateDailyRevenueRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var dr smartcreate.DailyRevenueRecord
	err := c.ShouldBindJSON(&dr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = drService.CreateDailyRevenueRecord(ctx, &dr)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteDailyRevenueRecord 删除日收入统计
// @Tags DailyRevenueRecord
// @Summary 删除日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.DailyRevenueRecord true "删除日收入统计"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dr/deleteDailyRevenueRecord [delete]
func (drApi *DailyRevenueRecordApi) DeleteDailyRevenueRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := drService.DeleteDailyRevenueRecord(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDailyRevenueRecordByIds 批量删除日收入统计
// @Tags DailyRevenueRecord
// @Summary 批量删除日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dr/deleteDailyRevenueRecordByIds [delete]
func (drApi *DailyRevenueRecordApi) DeleteDailyRevenueRecordByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := drService.DeleteDailyRevenueRecordByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDailyRevenueRecord 更新日收入统计
// @Tags DailyRevenueRecord
// @Summary 更新日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.DailyRevenueRecord true "更新日收入统计"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dr/updateDailyRevenueRecord [put]
func (drApi *DailyRevenueRecordApi) UpdateDailyRevenueRecord(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var dr smartcreate.DailyRevenueRecord
	err := c.ShouldBindJSON(&dr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = drService.UpdateDailyRevenueRecord(ctx, dr)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDailyRevenueRecord 用id查询日收入统计
// @Tags DailyRevenueRecord
// @Summary 用id查询日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询日收入统计"
// @Success 200 {object} response.Response{data=smartcreate.DailyRevenueRecord,msg=string} "查询成功"
// @Router /dr/findDailyRevenueRecord [get]
func (drApi *DailyRevenueRecordApi) FindDailyRevenueRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	redr, err := drService.GetDailyRevenueRecord(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(redr, c)
}

// GetDailyRevenueRecordList 分页获取日收入统计列表
// @Tags DailyRevenueRecord
// @Summary 分页获取日收入统计列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.DailyRevenueRecordSearch true "分页获取日收入统计列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dr/getDailyRevenueRecordList [get]
func (drApi *DailyRevenueRecordApi) GetDailyRevenueRecordList(c *gin.Context) {
	drService.InstallDailyRevenueRecord()
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo smartcreateReq.DailyRevenueRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := drService.GetDailyRevenueRecordInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetDailyRevenueRecordPublic 不需要鉴权的日收入统计接口
// @Tags DailyRevenueRecord
// @Summary 不需要鉴权的日收入统计接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dr/getDailyRevenueRecordPublic [get]
func (drApi *DailyRevenueRecordApi) GetDailyRevenueRecordPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	drService.GetDailyRevenueRecordPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的日收入统计接口信息",
	}, "获取成功", c)
}
