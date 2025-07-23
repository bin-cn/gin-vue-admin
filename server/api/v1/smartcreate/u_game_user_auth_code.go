package smartcreate

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
    smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GameUserAuthCodeApi struct {}



// CreateGameUserAuthCode 创建用户授权码码登录信息
// @Tags GameUserAuthCode
// @Summary 创建用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameUserAuthCode true "创建用户授权码码登录信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /game_user_auth_code/createGameUserAuthCode [post]
func (game_user_auth_codeApi *GameUserAuthCodeApi) CreateGameUserAuthCode(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var game_user_auth_code smartcreate.GameUserAuthCode
	err := c.ShouldBindJSON(&game_user_auth_code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = game_user_auth_codeService.CreateGameUserAuthCode(ctx,&game_user_auth_code)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteGameUserAuthCode 删除用户授权码码登录信息
// @Tags GameUserAuthCode
// @Summary 删除用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameUserAuthCode true "删除用户授权码码登录信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /game_user_auth_code/deleteGameUserAuthCode [delete]
func (game_user_auth_codeApi *GameUserAuthCodeApi) DeleteGameUserAuthCode(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := game_user_auth_codeService.DeleteGameUserAuthCode(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGameUserAuthCodeByIds 批量删除用户授权码码登录信息
// @Tags GameUserAuthCode
// @Summary 批量删除用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /game_user_auth_code/deleteGameUserAuthCodeByIds [delete]
func (game_user_auth_codeApi *GameUserAuthCodeApi) DeleteGameUserAuthCodeByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := game_user_auth_codeService.DeleteGameUserAuthCodeByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGameUserAuthCode 更新用户授权码码登录信息
// @Tags GameUserAuthCode
// @Summary 更新用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameUserAuthCode true "更新用户授权码码登录信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /game_user_auth_code/updateGameUserAuthCode [put]
func (game_user_auth_codeApi *GameUserAuthCodeApi) UpdateGameUserAuthCode(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var game_user_auth_code smartcreate.GameUserAuthCode
	err := c.ShouldBindJSON(&game_user_auth_code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = game_user_auth_codeService.UpdateGameUserAuthCode(ctx,game_user_auth_code)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGameUserAuthCode 用id查询用户授权码码登录信息
// @Tags GameUserAuthCode
// @Summary 用id查询用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户授权码码登录信息"
// @Success 200 {object} response.Response{data=smartcreate.GameUserAuthCode,msg=string} "查询成功"
// @Router /game_user_auth_code/findGameUserAuthCode [get]
func (game_user_auth_codeApi *GameUserAuthCodeApi) FindGameUserAuthCode(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	regame_user_auth_code, err := game_user_auth_codeService.GetGameUserAuthCode(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(regame_user_auth_code, c)
}
// GetGameUserAuthCodeList 分页获取用户授权码码登录信息列表
// @Tags GameUserAuthCode
// @Summary 分页获取用户授权码码登录信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.GameUserAuthCodeSearch true "分页获取用户授权码码登录信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /game_user_auth_code/getGameUserAuthCodeList [get]
func (game_user_auth_codeApi *GameUserAuthCodeApi) GetGameUserAuthCodeList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo smartcreateReq.GameUserAuthCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := game_user_auth_codeService.GetGameUserAuthCodeInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetGameUserAuthCodePublic 不需要鉴权的用户授权码码登录信息接口
// @Tags GameUserAuthCode
// @Summary 不需要鉴权的用户授权码码登录信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /game_user_auth_code/getGameUserAuthCodePublic [get]
func (game_user_auth_codeApi *GameUserAuthCodeApi) GetGameUserAuthCodePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    game_user_auth_codeService.GetGameUserAuthCodePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的用户授权码码登录信息接口信息",
    }, "获取成功", c)
}
