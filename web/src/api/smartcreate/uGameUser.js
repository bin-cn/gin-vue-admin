import service from '@/utils/request'
// @Tags GameUser
// @Summary 创建用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameUser true "创建用户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /game_user/createGameUser [post]
export const createGameUser = (data) => {
  return service({
    url: '/game_user/createGameUser',
    method: 'post',
    data
  })
}

// @Tags GameUser
// @Summary 删除用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameUser true "删除用户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /game_user/deleteGameUser [delete]
export const deleteGameUser = (params) => {
  return service({
    url: '/game_user/deleteGameUser',
    method: 'delete',
    params
  })
}

// @Tags GameUser
// @Summary 批量删除用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /game_user/deleteGameUser [delete]
export const deleteGameUserByIds = (params) => {
  return service({
    url: '/game_user/deleteGameUserByIds',
    method: 'delete',
    params
  })
}

// @Tags GameUser
// @Summary 更新用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameUser true "更新用户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /game_user/updateGameUser [put]
export const updateGameUser = (data) => {
  return service({
    url: '/game_user/updateGameUser',
    method: 'put',
    data
  })
}

// @Tags GameUser
// @Summary 用id查询用户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.GameUser true "用id查询用户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /game_user/findGameUser [get]
export const findGameUser = (params) => {
  return service({
    url: '/game_user/findGameUser',
    method: 'get',
    params
  })
}

// @Tags GameUser
// @Summary 分页获取用户信息表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /game_user/getGameUserList [get]
export const getGameUserList = (params) => {
  return service({
    url: '/game_user/getGameUserList',
    method: 'get',
    params
  })
}
// @Tags GameUser
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /game_user/findGameUserDataSource [get]
export const getGameUserDataSource = () => {
  return service({
    url: '/game_user/getGameUserDataSource',
    method: 'get',
  })
}

// @Tags GameUser
// @Summary 不需要鉴权的用户信息表接口
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.GameUserSearch true "分页获取用户信息表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /game_user/getGameUserPublic [get]
export const getGameUserPublic = () => {
  return service({
    url: '/game_user/getGameUserPublic',
    method: 'get',
  })
}
// SubmitGoldCoinInfo 提交用户资产信息
// @Tags GameUser
// @Summary 提交用户资产信息
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /game_user/submit_gold_coin_info [POST]
export const submit_gold_coin_info = () => {
  return service({
    url: '/game_user/submit_gold_coin_info',
    method: 'POST'
  })
}
