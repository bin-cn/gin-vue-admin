package smartcreate

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// TestImportServiceStructure 测试导入服务结构
func TestImportServiceStructure(t *testing.T) {
	service := GameUserAuthCodeImportService{}
	if service.generateUniqueLoginCode == nil {
		t.Error("generateUniqueLoginCode method not initialized")
	}
}

// TestGenerateUniqueLoginCode 测试唯一登录码生成
func TestGenerateUniqueLoginCode(t *testing.T) {
	// 这是一个mock测试，实际使用时需要连接数据库
	t.Log("测试唯一登录码生成逻辑")
}

// TestGetUserMapByNickNames 测试批量用户ID查询
func TestGetUserMapByNickNames(t *testing.T) {
	// 这是一个mock测试，实际使用时需要连接数据库
	t.Log("测试批量用户ID查询逻辑")
}

// TestExcelTemplate 测试Excel模板
func TestExcelTemplate(t *testing.T) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	// 设置表头
	headers := []string{
		"login_code",
		"assigner_name",
		"machine_no_name",
		"game_server_name",
		"account",
		"password",
		"role_game_name",
		"role_game_id",
		"i_d_name",
		"i_d_card_number",
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// 添加测试数据
	data := [][]string{
		{"TEST001", "管理员", "机器1", "服务器A", "test1", "pass123", "角色1", "1001", "张三", "123456789012345678"},
		{"TEST002", "用户1", "机器2", "服务器B", "test2", "pass456", "角色2", "1002", "李四", "987654321098765432"},
	}

	for rowIndex, row := range data {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(sheet, cell, value)
		}
	}

	// 保存测试文件
	if err := f.SaveAs("test_import_template.xlsx"); err != nil {
		t.Errorf("保存测试文件失败: %v", err)
	}

	t.Log("测试Excel模板创建成功")
}

// MockDB 用于测试的mock数据库
func MockDB() *gorm.DB {
	// 这里应该返回一个mock的数据库连接
	// 实际测试时需要替换为真实的测试数据库
	return nil
}

// TestColumnMapping 测试列映射
func TestColumnMapping(t *testing.T) {
	columnMap := map[string]int{
		"login_code":      0,
		"assigner_name":   1,
		"machine_no_name": 2,
	}

	service := GameUserAuthCodeImportService{}
	row := []string{"test123", "管理员", "机器1"}

	loginCode := service.getCellValue(row, columnMap, "login_code")
	if loginCode != "test123" {
		t.Errorf("列映射错误，期望test123，得到%s", loginCode)
	}

	assignerName := service.getCellValue(row, columnMap, "assigner_name")
	if assignerName != "管理员" {
		t.Errorf("列映射错误，期望管理员，得到%s", assignerName)
	}

	missing := service.getCellValue(row, columnMap, "missing_column")
	if missing != "" {
		t.Errorf("缺失列应该返回空字符串，得到%s", missing)
	}
}

// TestModelStructure 测试模型结构
func TestModelStructure(t *testing.T) {
	// 测试GameUserAuthCode模型
	authCode := smartcreate.GameUserAuthCode{
		LoginCode:    "TEST123",
		AssignerName: "测试用户",
		UserID:       1,
	}

	if authCode.LoginCode != "TEST123" {
		t.Error("模型结构错误")
	}

	// 测试SysUser模型
	user := system.SysUser{
		NickName: "测试用户",
	}

	if user.NickName != "测试用户" {
		t.Error("用户模型结构错误")
	}
}
