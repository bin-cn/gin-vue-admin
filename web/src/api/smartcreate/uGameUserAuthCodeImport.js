import service from '@/utils/request'

// @Tags GameUserAuthCode
// @Summary 自定义Excel导入（处理登录码唯一性和用户验证）
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce application/json
// @Param file formData file true "Excel文件"
// @Success 200 {object} response.Response{msg=string} "导入成功"
// @Router /game_user_auth_code/importCustom [post]
export const importGameUserAuthCodeCustom = (file) => {
  const formData = new FormData()
  formData.append('file', file)
  
  return service({
    url: '/game_user_auth_code/importCustom',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

// @Tags GameUserAuthCode
// @Summary 下载Excel导入模板
// @Security ApiKeyAuth
// @Produce application/octet-stream
// @Success 200 {file} file "模板文件"
// @Router /game_user_auth_code/downloadTemplate [get]
export const downloadGameUserAuthCodeTemplate = () => {
  return service({
    url: '/game_user_auth_code/downloadTemplate',
    method: 'get',
    responseType: 'blob'
  })
}