package smartcreate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service/smartcreate"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type GameUserAuthCodeImportApi struct{}

// ImportExcelWithCustomLogic 自定义Excel导入接口
// @Tags GameUserAuthCode
// @Summary 自定义Excel导入（处理登录码唯一性和用户验证）
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce application/json
// @Param file formData file true "Excel文件"
// @Success 200 {object} response.Response{msg=string} "导入成功"// @Router /game_user_auth_code/importCustom [post]
func (api *GameUserAuthCodeImportApi) ImportExcelWithCustomLogic(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("获取上传文件失败", zap.Error(err))
		response.FailWithMessage("获取上传文件失败", c)
		return
	}

	// 检查文件类型
	if !utils.CheckFileExt(file.Filename, []string{".xlsx", ".xls"}) {
		response.FailWithMessage("请上传Excel文件(.xlsx或.xls)", c)
		return
	}

	// 保存上传的文件到临时目录
	tempDir := os.TempDir()
	tempFile := filepath.Join(tempDir, file.Filename)
	if err := c.SaveUploadedFile(file, tempFile); err != nil {
		global.GVA_LOG.Error("保存上传文件失败", zap.Error(err))
		response.FailWithMessage("保存上传文件失败", c)
		return
	}
	defer os.Remove(tempFile)

	// 打开Excel文件
	excelFile, err := excelize.OpenFile(tempFile)
	if err != nil {
		global.GVA_LOG.Error("打开Excel文件失败", zap.Error(err))
		response.FailWithMessage("打开Excel文件失败: "+err.Error(), c)
		return
	}
	defer excelFile.Close()

	// 调用自定义导入服务
	importService := smartcreate.GameUserAuthCodeImportService{}
	err = importService.ImportExcelWithCustomLogic(c.Request.Context(), excelFile)
	if err != nil {
		global.GVA_LOG.Error("导入数据失败", zap.Error(err))
		response.FailWithMessage("导入数据失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("导入成功", c)
}

// DownloadImportTemplate 下载导入模板
// @Tags GameUserAuthCode
// @Summary 下载Excel导入模板
// @Security ApiKeyAuth
// @Produce application/octet-stream
// @Success 200 {file} file "模板文件"
// @Router /game_user_auth_code/downloadTemplate [get]
func (api *GameUserAuthCodeImportApi) DownloadImportTemplate(c *gin.Context) {
	// 创建Excel模板
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	// 设置表头
	headers := []string{
		"登录码",
		"使用人",
		"机器码名称",
		"区服",
		"账号",
		"密码",
		"游戏角色名字",
		"角色游戏ID",
		"ID名字",
		"身份证号码",
		"区服ID",
		"开服时间",
		"进服时间",
		"备注",
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// 设置响应头
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=game_user_auth_code_template.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	// 写入响应
	if err := f.Write(c.Writer); err != nil {
		global.GVA_LOG.Error("生成模板文件失败", zap.Error(err))
		response.FailWithMessage("生成模板文件失败", c)
		return
	}
}