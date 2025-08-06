import service from '@/utils/request'

// @Tags SysTaskLog
// @Summary 创建sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysTaskLog true "创建sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysTaskLog/createSysTaskLog [post]
export const createSysTaskLog = (data) => {
  return service({
    url: '/sysTaskLog/createSysTaskLog',
    method: 'post',
    data
  })
}

// @Tags SysTaskLog
// @Summary 删除sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysTaskLog true "删除sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysTaskLog/deleteSysTaskLog [delete]
export const deleteSysTaskLog = (params) => {
  return service({
    url: '/sysTaskLog/deleteSysTaskLog',
    method: 'delete',
    params
  })
}

// @Tags SysTaskLog
// @Summary 批量删除sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysTaskLog/deleteSysTaskLog [delete]
export const deleteSysTaskLogByIds = (params) => {
  return service({
    url: '/sysTaskLog/deleteSysTaskLogByIds',
    method: 'delete',
    params
  })
}

// @Tags SysTaskLog
// @Summary 更新sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysTaskLog true "更新sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysTaskLog/updateSysTaskLog [put]
export const updateSysTaskLog = (data) => {
  return service({
    url: '/sysTaskLog/updateSysTaskLog',
    method: 'put',
    data
  })
}

// @Tags SysTaskLog
// @Summary 用id查询sysTaskLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysTaskLog true "用id查询sysTaskLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysTaskLog/findSysTaskLog [get]
export const findSysTaskLog = (params) => {
  return service({
    url: '/sysTaskLog/findSysTaskLog',
    method: 'get',
    params
  })
}

// @Tags SysTaskLog
// @Summary 分页获取sysTaskLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysTaskLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysTaskLog/getSysTaskLogList [get]
export const getSysTaskLogList = (params) => {
  return service({
    url: '/sysTaskLog/getSysTaskLogList',
    method: 'get',
    params
  })
}
