import service from '@/utils/request'
// @Tags DailyRevenueRecord
// @Summary 创建日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DailyRevenueRecord true "创建日收入统计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dr/createDailyRevenueRecord [post]
export const createDailyRevenueRecord = (data) => {
  return service({
    url: '/dr/createDailyRevenueRecord',
    method: 'post',
    data
  })
}

// @Tags DailyRevenueRecord
// @Summary 删除日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DailyRevenueRecord true "删除日收入统计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dr/deleteDailyRevenueRecord [delete]
export const deleteDailyRevenueRecord = (params) => {
  return service({
    url: '/dr/deleteDailyRevenueRecord',
    method: 'delete',
    params
  })
}

// @Tags DailyRevenueRecord
// @Summary 批量删除日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除日收入统计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dr/deleteDailyRevenueRecord [delete]
export const deleteDailyRevenueRecordByIds = (params) => {
  return service({
    url: '/dr/deleteDailyRevenueRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags DailyRevenueRecord
// @Summary 更新日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DailyRevenueRecord true "更新日收入统计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dr/updateDailyRevenueRecord [put]
export const updateDailyRevenueRecord = (data) => {
  return service({
    url: '/dr/updateDailyRevenueRecord',
    method: 'put',
    data
  })
}

// @Tags DailyRevenueRecord
// @Summary 用id查询日收入统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DailyRevenueRecord true "用id查询日收入统计"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dr/findDailyRevenueRecord [get]
export const findDailyRevenueRecord = (params) => {
  return service({
    url: '/dr/findDailyRevenueRecord',
    method: 'get',
    params
  })
}

// @Tags DailyRevenueRecord
// @Summary 分页获取日收入统计列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取日收入统计列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dr/getDailyRevenueRecordList [get]
export const getDailyRevenueRecordList = (params) => {
  return service({
    url: '/dr/getDailyRevenueRecordList',
    method: 'get',
    params
  })
}


export const getStatistic = (params) => {
  return service({
    url: '/dr/getStatistic',
    method: 'get',
    params
  })
}


// @Tags DailyRevenueRecord
// @Summary 不需要鉴权的日收入统计接口
// @Accept application/json
// @Produce application/json
// @Param data query smartcreateReq.DailyRevenueRecordSearch true "分页获取日收入统计列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dr/getDailyRevenueRecordPublic [get]
export const getDailyRevenueRecordPublic = () => {
  return service({
    url: '/dr/getDailyRevenueRecordPublic',
    method: 'get',
  })
}
