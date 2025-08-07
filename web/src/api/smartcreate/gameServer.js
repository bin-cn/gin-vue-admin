import service from '@/utils/request'
// @Tags GameServer
// @Summary 创建游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameServer true "创建游戏服务器列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /gamesrv/createGameServer [post]
export const createGameServer = (data) => {
  return service({
    url: '/gamesrv/createGameServer',
    method: 'post',
    data
  })
}

// @Tags GameServer
// @Summary 删除游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameServer true "删除游戏服务器列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gamesrv/deleteGameServer [delete]
export const deleteGameServer = (params) => {
  return service({
    url: '/gamesrv/deleteGameServer',
    method: 'delete',
    params
  })
}

// @Tags GameServer
// @Summary 批量删除游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除游戏服务器列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gamesrv/deleteGameServer [delete]
export const deleteGameServerByIds = (params) => {
  return service({
    url: '/gamesrv/deleteGameServerByIds',
    method: 'delete',
    params
  })
}

// @Tags GameServer
// @Summary 更新游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameServer true "更新游戏服务器列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gamesrv/updateGameServer [put]
export const updateGameServer = (data) => {
  return service({
    url: '/gamesrv/updateGameServer',
    method: 'put',
    data
  })
}

// @Tags GameServer
// @Summary 用id查询游戏服务器列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.GameServer true "用id查询游戏服务器列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gamesrv/findGameServer [get]
export const findGameServer = (params) => {
  return service({
    url: '/gamesrv/findGameServer',
    method: 'get',
    params
  })
}

// @Tags GameServer
// @Summary 分页获取游戏服务器列表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取游戏服务器列表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gamesrv/getGameServerList [get]
export const getGameServerList = (params) => {
  return service({
    url: '/gamesrv/getGameServerList',
    method: 'get',
    params
  })
}

// @Tags GameServer
// @Summary 不需要鉴权的游戏服务器列表接口
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.GameServerSearch true "分页获取游戏服务器列表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /gamesrv/getGameServerPublic [get]
export const getGameServerPublic = () => {
  return service({
    url: '/gamesrv/getGameServerPublic',
    method: 'get',
  })
}
