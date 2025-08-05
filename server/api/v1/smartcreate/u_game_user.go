package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GameUserApi struct{}

// CreateGameUser 创建用户信息表
// @Tags GameUser
// @Summary 创建用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameUser true "创建用户信息表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /game_user/createGameUser [post]
func (game_userApi *GameUserApi) CreateGameUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var game_user smartcreate.GameUser
	err := c.ShouldBindJSON(&game_user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	game_user.CreatedBy = utils.GetUserID(c)
	err = game_userService.CreateGameUser(ctx, &game_user)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteGameUser 删除用户信息表
// @Tags GameUser
// @Summary 删除用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameUser true "删除用户信息表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /game_user/deleteGameUser [delete]
func (game_userApi *GameUserApi) DeleteGameUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := game_userService.DeleteGameUser(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGameUserByIds 批量删除用户信息表
// @Tags GameUser
// @Summary 批量删除用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /game_user/deleteGameUserByIds [delete]
func (game_userApi *GameUserApi) DeleteGameUserByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := game_userService.DeleteGameUserByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGameUser 更新用户信息表
// @Tags GameUser
// @Summary 更新用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameUser true "更新用户信息表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /game_user/updateGameUser [put]
func (game_userApi *GameUserApi) UpdateGameUser(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var game_user smartcreate.GameUser
	err := c.ShouldBindJSON(&game_user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	game_user.UpdatedBy = utils.GetUserID(c)
	err = game_userService.UpdateGameUser(ctx, game_user)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGameUser 用id查询用户信息表
// @Tags GameUser
// @Summary 用id查询用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户信息表"
// @Success 200 {object} response.Response{data=smartcreate.GameUser,msg=string} "查询成功"
// @Router /game_user/findGameUser [get]
func (game_userApi *GameUserApi) FindGameUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	regame_user, err := game_userService.GetGameUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(regame_user, c)
}

// GetGameUserList 分页获取用户信息表列表
// @Tags GameUser
// @Summary 分页获取用户信息表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.GameUserSearch true "分页获取用户信息表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /game_user/getGameUserList [get]
func (game_userApi *GameUserApi) GetGameUserList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo smartcreateReq.GameUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := game_userService.GetGameUserInfoList(ctx, pageInfo)
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

// GetGameUserDataSource 获取GameUser的数据源
// @Tags GameUser
// @Summary 获取GameUser的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /game_user/getGameUserDataSource [get]
func (game_userApi *GameUserApi) GetGameUserDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := game_userService.GetGameUserDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetGameUserPublic 不需要鉴权的用户信息表接口
// @Tags GameUser
// @Summary 不需要鉴权的用户信息表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /game_user/getGameUserPublic [get]
func (game_userApi *GameUserApi) GetGameUserPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	game_userService.GetGameUserPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户信息表接口信息",
	}, "获取成功", c)
}

// SubmitGoldCoinInfo 提交用户资产信息
// @Tags GameUser
// @Summary 提交用户资产信息
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.GameUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /game_user/submit_gold_coin_info [POST]
func (game_userApi *GameUserApi) SubmitGoldCoinInfo(c *gin.Context) {

	var Submitinfo smartcreateReq.SubmitGoldCoinInfo

	err := c.ShouldBindJSON(&Submitinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Submitinfo.UpdatedBy = utils.GetUserID(c)
	// err = game_userService.UpdateGameUser(c,game_user)

	// 创建业务用Context
	ctx := c.Request.Context()
	// 请添加自己的业务逻辑
	err = game_userService.SubmitGoldCoinInfo(ctx, Submitinfo)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}
