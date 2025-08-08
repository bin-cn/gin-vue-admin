import service from '@/utils/request'
// @Tags CookieData
// @Summary 创建网站cookie
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CookieData true "创建网站cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wc/createCookieData [post]
export const createCookieData = (data) => {
  return service({
    url: '/wc/createCookieData',
    method: 'post',
    data
  })
}

// @Tags CookieData
// @Summary 删除网站cookie
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CookieData true "删除网站cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wc/deleteCookieData [delete]
export const deleteCookieData = (params) => {
  return service({
    url: '/wc/deleteCookieData',
    method: 'delete',
    params
  })
}

// @Tags CookieData
// @Summary 批量删除网站cookie
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除网站cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wc/deleteCookieData [delete]
export const deleteCookieDataByIds = (params) => {
  return service({
    url: '/wc/deleteCookieDataByIds',
    method: 'delete',
    params
  })
}

// @Tags CookieData
// @Summary 更新网站cookie
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CookieData true "更新网站cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wc/updateCookieData [put]
export const updateCookieData = (data) => {
  return service({
    url: '/wc/updateCookieData',
    method: 'put',
    data
  })
}

// @Tags CookieData
// @Summary 用id查询网站cookie
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.CookieData true "用id查询网站cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wc/findCookieData [get]
export const findCookieData = (params) => {
  return service({
    url: '/wc/findCookieData',
    method: 'get',
    params
  })
}

// @Tags CookieData
// @Summary 分页获取网站cookie列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取网站cookie列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wc/getCookieDataList [get]
export const getCookieDataList = (params) => {
  return service({
    url: '/wc/getCookieDataList',
    method: 'get',
    params
  })
}

// @Tags CookieData
// @Summary 不需要鉴权的网站cookie接口
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.CookieDataSearch true "分页获取网站cookie列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wc/getCookieDataPublic [get]
export const getCookieDataPublic = () => {
  return service({
    url: '/wc/getCookieDataPublic',
    method: 'get',
  })
}
