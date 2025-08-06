import service from '@/utils/request'

// @Tags SysTask
// @Summary 创建sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysTask true "创建sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysTask/createSysTask [post]
export const createSysTask = (data) => {
  return service({
    url: '/sysTask/createSysTask',
    method: 'post',
    data
  })
}

// @Tags SysTask
// @Summary 删除sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysTask true "删除sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysTask/deleteSysTask [delete]
export const deleteSysTask = (params) => {
  return service({
    url: '/sysTask/deleteSysTask',
    method: 'delete',
    params
  })
}

// @Tags SysTask
// @Summary 批量删除sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysTask/deleteSysTask [delete]
export const deleteSysTaskByIds = (params) => {
  return service({
    url: '/sysTask/deleteSysTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags SysTask
// @Summary 更新sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysTask true "更新sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysTask/updateSysTask [put]
export const updateSysTask = (data) => {
  return service({
    url: '/sysTask/updateSysTask',
    method: 'put',
    data
  })
}

// @Tags SysTask
// @Summary 用id查询sysTask表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysTask true "用id查询sysTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysTask/findSysTask [get]
export const findSysTask = (params) => {
  return service({
    url: '/sysTask/findSysTask',
    method: 'get',
    params
  })
}

// @Tags SysTask
// @Summary 分页获取sysTask表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysTask表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysTask/getSysTaskList [get]
export const getSysTaskList = (params) => {
  return service({
    url: '/sysTask/getSysTaskList',
    method: 'get',
    params
  })
}

// 运行一次
export const runOneSysTask = (params) => {
  return service({
    url: '/sysTask/runOneSysTask',
    method: 'get',
    params
  })
}

// 激活任务
export const enableSysTask = (params) => {
  return service({
    url: '/sysTask/enableSysTask',
    method: 'get',
    params
  })
}

// 暂停任务
export const disableSysTask = (params) => {
  return service({
    url: '/sysTask/disableSysTask',
    method: 'get',
    params
  })
}

// 删除任务
export const removeSysTask = (params) => {
  return service({
    url: '/sysTask/removeSysTask',
    method: 'get',
    params
  })
}
