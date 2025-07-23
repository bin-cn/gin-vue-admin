import service from '@/utils/request'
// @Tags GameUserAuthCode
// @Summary 创建用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameUserAuthCode true "创建用户授权码码登录信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /game_user_auth_code/createGameUserAuthCode [post]
export const createGameUserAuthCode = (data) => {
  return service({
    url: '/game_user_auth_code/createGameUserAuthCode',
    method: 'post',
    data
  })
}

// @Tags GameUserAuthCode
// @Summary 删除用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameUserAuthCode true "删除用户授权码码登录信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /game_user_auth_code/deleteGameUserAuthCode [delete]
export const deleteGameUserAuthCode = (params) => {
  return service({
    url: '/game_user_auth_code/deleteGameUserAuthCode',
    method: 'delete',
    params
  })
}

// @Tags GameUserAuthCode
// @Summary 批量删除用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户授权码码登录信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /game_user_auth_code/deleteGameUserAuthCode [delete]
export const deleteGameUserAuthCodeByIds = (params) => {
  return service({
    url: '/game_user_auth_code/deleteGameUserAuthCodeByIds',
    method: 'delete',
    params
  })
}

// @Tags GameUserAuthCode
// @Summary 更新用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GameUserAuthCode true "更新用户授权码码登录信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /game_user_auth_code/updateGameUserAuthCode [put]
export const updateGameUserAuthCode = (data) => {
  return service({
    url: '/game_user_auth_code/updateGameUserAuthCode',
    method: 'put',
    data
  })
}

// @Tags GameUserAuthCode
// @Summary 用id查询用户授权码码登录信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.GameUserAuthCode true "用id查询用户授权码码登录信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /game_user_auth_code/findGameUserAuthCode [get]
export const findGameUserAuthCode = (params) => {
  return service({
    url: '/game_user_auth_code/findGameUserAuthCode',
    method: 'get',
    params
  })
}

// @Tags GameUserAuthCode
// @Summary 分页获取用户授权码码登录信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户授权码码登录信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /game_user_auth_code/getGameUserAuthCodeList [get]
export const getGameUserAuthCodeList = (params) => {
  return service({
    url: '/game_user_auth_code/getGameUserAuthCodeList',
    method: 'get',
    params
  })
}

// @Tags GameUserAuthCode
// @Summary 不需要鉴权的用户授权码码登录信息接口
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.GameUserAuthCodeSearch true "分页获取用户授权码码登录信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /game_user_auth_code/getGameUserAuthCodePublic [get]
export const getGameUserAuthCodePublic = () => {
  return service({
    url: '/game_user_auth_code/getGameUserAuthCodePublic',
    method: 'get',
  })
}
