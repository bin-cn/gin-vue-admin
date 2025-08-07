package smartcreate

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
    smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GameServerApi struct {}



// CreateGameServer 创建游戏服务器列表
// @Tags GameServer
// @Summary 创建游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameServer true "创建游戏服务器列表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /gamesrv/createGameServer [post]
func (gamesrvApi *GameServerApi) CreateGameServer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var gamesrv smartcreate.GameServer
	err := c.ShouldBindJSON(&gamesrv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = gamesrvService.CreateGameServer(ctx,&gamesrv)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteGameServer 删除游戏服务器列表
// @Tags GameServer
// @Summary 删除游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameServer true "删除游戏服务器列表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /gamesrv/deleteGameServer [delete]
func (gamesrvApi *GameServerApi) DeleteGameServer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := gamesrvService.DeleteGameServer(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGameServerByIds 批量删除游戏服务器列表
// @Tags GameServer
// @Summary 批量删除游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /gamesrv/deleteGameServerByIds [delete]
func (gamesrvApi *GameServerApi) DeleteGameServerByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := gamesrvService.DeleteGameServerByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGameServer 更新游戏服务器列表
// @Tags GameServer
// @Summary 更新游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body smartcreate.GameServer true "更新游戏服务器列表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /gamesrv/updateGameServer [put]
func (gamesrvApi *GameServerApi) UpdateGameServer(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var gamesrv smartcreate.GameServer
	err := c.ShouldBindJSON(&gamesrv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = gamesrvService.UpdateGameServer(ctx,gamesrv)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGameServer 用id查询游戏服务器列表
// @Tags GameServer
// @Summary 用id查询游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询游戏服务器列表"
// @Success 200 {object} response.Response{data=smartcreate.GameServer,msg=string} "查询成功"
// @Router /gamesrv/findGameServer [get]
func (gamesrvApi *GameServerApi) FindGameServer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	regamesrv, err := gamesrvService.GetGameServer(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(regamesrv, c)
}
// GetGameServerList 分页获取游戏服务器列表列表
// @Tags GameServer
// @Summary 分页获取游戏服务器列表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.GameServerSearch true "分页获取游戏服务器列表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /gamesrv/getGameServerList [get]
func (gamesrvApi *GameServerApi) GetGameServerList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo smartcreateReq.GameServerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := gamesrvService.GetGameServerInfoList(ctx,pageInfo)
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

// GetGameServerPublic 不需要鉴权的游戏服务器列表接口
// @Tags GameServer
// @Summary 不需要鉴权的游戏服务器列表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /gamesrv/getGameServerPublic [get]
func (gamesrvApi *GameServerApi) GetGameServerPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    gamesrvService.GetGameServerPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的游戏服务器列表接口信息",
    }, "获取成功", c)
}
