package smartcreate

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

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
			return fmt.Errorf("Excel缺少必需列: %s", col)
		}
	}

	// 收集所有assigner_name用于批量查询
	assignerNames := make([]string, 0)
	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]
		if len(row) == 0 {
			continue
		}
		assignerName := s.getCellValue(row, columnMap, "使用人")
		if assignerName != "" {
			assignerNames = append(assignerNames, assignerName)
		}
	}

	// 开始事务
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 批量查询用户数据
		userMap, err := s.getUserMapByNickNames(tx, assignerNames)
		if err != nil {
			return fmt.Errorf("批量查询用户数据失败: %w", err)
		}

		// 准备批量插入的数据
		authCodes := make([]*smartcreate.GameUserAuthCode, 0)

		// 收集所有role_game_id用于批量检查唯一性
		roleGameIds := make([]string, 0)
		for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
			row := rows[rowIndex]
			if len(row) == 0 {
				continue
			}
			roleGameId := s.getCellValue(row, columnMap, "角色游戏ID")
			if roleGameId != "" {
				roleGameIds = append(roleGameIds, roleGameId)
			}
		}

		// 批量检查role_game_id的唯一性
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

		// 处理每一行数据
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

			// 检查登录码是否已存在
			uniqueLoginCode, err := s.generateUniqueLoginCode(tx, loginCode)
			if err != nil {
				return fmt.Errorf("第%d行: 生成唯一登录码失败: %w", rowIndex+1, err)
			}
			authCode.LoginCode = uniqueLoginCode

			// 处理使用人查询用户ID
			assignerName := s.getCellValue(row, columnMap, "使用人")
			if assignerName == "" {
				return fmt.Errorf("第%d行: 使用人不能为空", rowIndex+1)
			}
			authCode.AssignerName = &assignerName

			userID, exists := userMap[assignerName]
			if !exists {
				return fmt.Errorf("第%d行: 用户 '%s' 不存在", rowIndex+1, assignerName)
			}
			authCode.UserId = &userID

			// 处理必填字段
			machineNoName := s.getCellValue(row, columnMap, "机器编号")
			if machineNoName == "" {
				return fmt.Errorf("第%d行: 机器编号不能为空", rowIndex+1)
			}
			authCode.MachineNoName = &machineNoName

			account := s.getCellValue(row, columnMap, "账号")
			if account == "" {
				return fmt.Errorf("第%d行: 账号不能为空", rowIndex+1)
			}
			authCode.Account = &account

			password := s.getCellValue(row, columnMap, "密码")
			if password == "" {
				return fmt.Errorf("第%d行: 密码不能为空", rowIndex+1)
			}
			authCode.Password = &password

			gameServerName := s.getCellValue(row, columnMap, "区服")
			if gameServerName == "" {
				return fmt.Errorf("第%d行: 区服不能为空", rowIndex+1)
			}
			authCode.GameServerName = &gameServerName

			roleGameName := s.getCellValue(row, columnMap, "游戏角色名字")
			if roleGameName == "" {
				return fmt.Errorf("第%d行: 游戏角色名字不能为空", rowIndex+1)
			}
			authCode.RoleGameName = &roleGameName

			idName := s.getCellValue(row, columnMap, "ID名字")
			if idName == "" {
				return fmt.Errorf("第%d行: ID名字不能为空", rowIndex+1)
			}
			authCode.IDName = &idName

			idCardNumber := s.getCellValue(row, columnMap, "身份证号码")
			if idCardNumber == "" {
				return fmt.Errorf("第%d行: 身份证号码不能为空", rowIndex+1)
			}
			authCode.IDCardNumber = &idCardNumber

			// 处理可选字段
			roleGameId := s.getCellValue(row, columnMap, "角色游戏ID")
			if roleGameId != "" {
				authCode.RoleGameId = &roleGameId
			}

			gameServerId := s.parseInt(s.getCellValue(row, columnMap, "区服ID"))
			if gameServerId > 0 {
				authCode.GameServerId = &gameServerId
			}

			serverOpenTime := s.getCellValue(row, columnMap, "开服时间")
			if serverOpenTime != "" {
				authCode.ServerOpenTime = &serverOpenTime
			}

			enterServerTime := s.getCellValue(row, columnMap, "进服时间")
			if enterServerTime != "" {
				authCode.EnterServerTime = &enterServerTime
			}

			remark := s.getCellValue(row, columnMap, "备注")
			if remark != "" {
				authCode.Remark = &remark
			}

			// 设置默认值
			accountStatus := "InUse"
			authCode.AccountStatus = &accountStatus

			// 添加到批量插入列表
			authCodes = append(authCodes, authCode)
		}

		// 批量插入记录
		if len(authCodes) > 0 {
			if err := tx.CreateInBatches(authCodes, 100).Error; err != nil {
				return fmt.Errorf("批量插入记录失败: %w", err)
			}
		}

		return nil
	})

	return err
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
