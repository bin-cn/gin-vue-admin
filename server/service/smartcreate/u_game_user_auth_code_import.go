package smartcreate

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// ExcelSerialToSlash 把字符串形式的 Excel 序列日期 转成 2025/6/3 0:00
func ExcelSerialToSlash(str string) string {
	if str == "" {
		return ""
	}

	// 如果是 6/3/25 00:00 格式
	if strings.Contains(str, "/") {
		if t, err := time.Parse("1/2/06 15:04", str); err == nil {
			return t.Format("2006/1/2 15:04")
		}
	}

	// 保持原有逻辑
	return str
}

// ExcelStringToString 把字符串形式的 Excel 序列日期 转成年-月-日 时:分
func ExcelStringToString(s string) string {
	serial, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return s
	}

	// Excel的基准日期是1899-12-30
	const baseDate = "1899-12-30"
	baseTime, _ := time.Parse("2006-01-02", baseDate)

	// 计算天数和当天的时间部分
	days := int64(serial)
	fraction := serial - float64(days)

	// 计算完整时间
	fullTime := baseTime.AddDate(0, 0, int(days))
	secondsInDay := int64(fraction * 24 * 60 * 60)
	resultTime := fullTime.Add(time.Duration(secondsInDay) * time.Second)

	return resultTime.Format("2006-01-02 15:04")
}

type GameUserAuthCodeImportService struct{}

// ImportExcelWithCustomLogic 自定义Excel导入，处理登录码唯一性和用户ID查询
func (s *GameUserAuthCodeImportService) ImportExcelWithCustomLogic(ctx context.Context, file *excelize.File) error {
	// 获取第一个工作表
	sheetName := file.GetSheetName(0)
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("读取Excel失败: %w", err)
	}

	if len(rows) < 2 {
		return fmt.Errorf("Excel文件格式不正确，至少需要标题行和数据行")
	}

	// 获取标题行并建立列索引映射
	headers := rows[0]
	columnMap := make(map[string]int)
	for i, header := range headers {
		columnMap[header] = i
	}

	// 检查必需的列
	requiredColumns := []string{"登录码", "使用人"}
	for _, col := range requiredColumns {
		if _, exists := columnMap[col]; !exists {
			return fmt.Errorf("excel缺少必需列: %s", col)
		}
	}

	// 收集所有assigner_name用于批量查询
	assignerNames := make([]string, 0)
	roleGameIds := make([]string, 0)

	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]
		if len(row) == 0 {
			continue
		}

		assignerName := s.getCellValue(row, columnMap, "使用人")
		if assignerName != "" {
			assignerNames = append(assignerNames, assignerName)
		}

		roleGameId := s.getCellValue(row, columnMap, "游戏角色ID")
		if roleGameId != "" {
			roleGameIds = append(roleGameIds, roleGameId)
		}
	}

	// 预处理所有数据
	authCodes := make([]*smartcreate.GameUserAuthCode, 0)

	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]
		if len(row) == 0 {
			continue // 跳过空行
		}

		// 创建实体
		authCode := &smartcreate.GameUserAuthCode{}

		// 处理登录码唯一性
		loginCode := s.getCellValue(row, columnMap, "登录码")
		if loginCode == "" {
			return fmt.Errorf("第%d行: 登录码不能为空", rowIndex+1)
		}

		// 处理必填字段验证
		assignerName := s.getCellValue(row, columnMap, "使用人")
		if assignerName == "" {
			return fmt.Errorf("第%d行: 使用人不能为空", rowIndex+1)
		}

		machineNoName := s.getCellValue(row, columnMap, "机器编号")
		if machineNoName == "" {
			return fmt.Errorf("第%d行: 机器编号不能为空", rowIndex+1)
		}

		account := s.getCellValue(row, columnMap, "账号")
		if account == "" {
			return fmt.Errorf("第%d行: 账号不能为空", rowIndex+1)
		}

		password := s.getCellValue(row, columnMap, "密码")
		if password == "" {
			return fmt.Errorf("第%d行: 密码不能为空", rowIndex+1)
		}

		gameServerName := s.getCellValue(row, columnMap, "区服")
		if gameServerName == "" {
			return fmt.Errorf("第%d行: 区服不能为空", rowIndex+1)
		}

		roleGameName := s.getCellValue(row, columnMap, "游戏角色名字")
		if roleGameName == "" {
			return fmt.Errorf("第%d行: 游戏角色名字不能为空", rowIndex+1)
		}

		idName := s.getCellValue(row, columnMap, "ID名字")
		if idName == "" {
			return fmt.Errorf("第%d行: ID名字不能为空", rowIndex+1)
		}

		idCardNumber := s.getCellValue(row, columnMap, "身份证号码")
		if idCardNumber == "" {
			return fmt.Errorf("第%d行: 身份证号码不能为空", rowIndex+1)
		}

		// 设置默认值
		now := time.Now()
		authCode.CreatedAt = now
		authCode.UpdatedAt = now
		authCode.LoginCode = &loginCode
		authCode.AssignerName = &assignerName
		authCode.MachineNoName = &machineNoName
		authCode.Account = &account
		authCode.Password = &password
		authCode.GameServerName = &gameServerName
		authCode.RoleGameName = &roleGameName
		authCode.IDName = &idName
		authCode.IDCardNumber = &idCardNumber

		// 处理可选字段
		roleGameId := s.getCellValue(row, columnMap, "游戏角色ID")
		if roleGameId != "" {
			authCode.RoleGameId = &roleGameId
		}

		gameServerId := s.parseInt(s.getCellValue(row, columnMap, "区服ID"))
		if gameServerId > 0 {
			authCode.GameServerId = &gameServerId
		}

		serverOpenTime := s.getCellValue(row, columnMap, "开区时间")
		if serverOpenTime != "" {
			serverOpenTime = ExcelSerialToSlash(serverOpenTime)
			authCode.ServerOpenTime = &serverOpenTime
		}

		enterServerTime := s.getCellValue(row, columnMap, "进区时间")
		if enterServerTime != "" {
			enterServerTime = ExcelStringToString(enterServerTime)
			authCode.EnterServerTime = &enterServerTime
		}

		remark := s.getCellValue(row, columnMap, "备注")
		if remark != "" {
			authCode.Remark = &remark
		}

		authCodes = append(authCodes, authCode)
	}

	// 开始事务，只包裹数据库操作
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 预加载game_server数据到map
		serverMap, err := s.loadGameServerMap(tx)
		if err != nil {
			return fmt.Errorf("预加载游戏服务器数据失败: %w", err)
		}

		// 批量查询用户数据
		userMap, err := s.getUserMapByNickNames(tx, assignerNames)
		if err != nil {
			return fmt.Errorf("批量查询用户数据失败: %w", err)
		}

		// 处理用户ID映射和登录码唯一性
		for _, authCode := range authCodes {
			// 处理用户ID
			userID, exists := userMap[*authCode.AssignerName]
			if !exists {
				return fmt.Errorf("用户 '%s' 不存在", *authCode.AssignerName)
			}
			authCode.UserId = &userID

			// 处理登录码唯一性
			uniqueLoginCode, err := s.generateUniqueLoginCode(tx, *authCode.LoginCode)
			if err != nil {
				return fmt.Errorf("生成唯一登录码失败: %w", err)
			}
			authCode.LoginCode = uniqueLoginCode

			// 处理区服ID映射
			if authCode.GameServerId != nil && *authCode.GameServerId > 0 {
				if serverInfo, exists := serverMap[*authCode.GameServerId]; exists && serverInfo.MainServerZoneId != nil {
					authCode.MainServerZone = serverInfo.MainServerZoneId
				}
			}
		}

		// 检查角色游戏ID唯一性
		if len(roleGameIds) > 0 {
			var existingRoleGameIds []string
			err = tx.Model(&smartcreate.GameUserAuthCode{}).
				Where("role_game_id IN ?", roleGameIds).
				Pluck("role_game_id", &existingRoleGameIds).Error
			if err != nil {
				return fmt.Errorf("检查角色游戏ID唯一性失败: %w", err)
			}
			if len(existingRoleGameIds) > 0 {
				return fmt.Errorf("以下角色游戏ID已存在: %v", existingRoleGameIds)
			}
		}

		// 批量创建记录 - 分批处理避免MySQL占位符限制
		const batchSize = 1000
		if len(authCodes) > 0 {
			totalRecords := len(authCodes)
			for i := 0; i < totalRecords; i += batchSize {
				end := i + batchSize
				if end > totalRecords {
					end = totalRecords
				}

				batch := authCodes[i:end]
				if err := tx.Create(&batch).Error; err != nil {
					return fmt.Errorf("批量创建授权码记录失败: %w", err)
				}
			}
		}

		return nil
	})
}

// loadGameServerMap 预加载game_server数据到map，提高效率
func (s *GameUserAuthCodeImportService) loadGameServerMap(tx *gorm.DB) (map[int]*smartcreate.GameServer, error) {
	var servers []smartcreate.GameServer
	err := tx.Find(&servers).Error
	if err != nil {
		return nil, err
	}

	serverMap := make(map[int]*smartcreate.GameServer)
	for i := range servers {
		server := &servers[i]
		if server.ServerZoneId != nil {
			// 将string类型的server_zone_id转换为int
			var zoneId int
			fmt.Sscanf(*server.ServerZoneId, "%d", &zoneId)
			serverMap[zoneId] = server
		}
	}

	return serverMap, nil
}

// generateUniqueLoginCode 生成唯一的登录码
func (s *GameUserAuthCodeImportService) generateUniqueLoginCode(tx *gorm.DB, originalCode string) (*string, error) {
	// 首先检查原始登录码是否已存在（包含软删除的数据）
	var count int64
	err := tx.Unscoped().Model(&smartcreate.GameUserAuthCode{}).Where("login_code = ?", originalCode).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &originalCode, nil // 原始登录码可用
	}

	// 生成新的唯一登录码
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ { // 最多尝试100次
		newCode := fmt.Sprintf("%s_%d", originalCode, rand.Intn(9999))
		var newCount int64
		err := tx.Unscoped().Model(&smartcreate.GameUserAuthCode{}).Where("login_code = ?", newCode).Count(&newCount).Error
		if err != nil {
			return nil, err
		}
		if newCount == 0 {
			return &newCode, nil
		}
	}

	return nil, fmt.Errorf("无法生成唯一登录码")
}

// getUserMapByNickNames 根据昵称批量查询用户ID
func (s *GameUserAuthCodeImportService) getUserMapByNickNames(tx *gorm.DB, nickNames []string) (map[string]int, error) {
	userMap := make(map[string]int)
	if len(nickNames) == 0 {
		return userMap, nil
	}

	var users []system.SysUser
	err := tx.Where("nick_name IN ?", nickNames).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	for _, user := range users {
		userMap[user.NickName] = int(user.ID)
	}

	return userMap, nil
}

func (s *GameUserAuthCodeImportService) getCellValueTime(row []string, columnMap map[string]int, columnName string) string {
	if index, exists := columnMap[columnName]; exists && index < len(row) {
		return strings.TrimSpace(row[index])
	}
	return ""
}

// getCellValue 安全获取单元格值
func (s *GameUserAuthCodeImportService) getCellValue(row []string, columnMap map[string]int, columnName string) string {
	if index, exists := columnMap[columnName]; exists && index < len(row) {
		return row[index]
	}
	return ""
}

// parseInt 安全解析整数
func (s *GameUserAuthCodeImportService) parseInt(value string) int {
	if value == "" {
		return 0
	}
	var result int
	fmt.Sscanf(value, "%d", &result)
	return result
}
