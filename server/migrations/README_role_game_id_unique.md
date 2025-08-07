# role_game_id字段唯一索引设置指南

## 概述
为game_user_auth_codes表的role_game_id字段添加唯一索引，确保角色游戏ID不会重复。

## 修改内容

### 1. 数据库修改
- **SQL文件**: `add_unique_index_role_game_id.sql`
- **操作**: 为role_game_id字段添加唯一索引
- **注意事项**: 执行前需确保没有重复值

### 2. 后端代码修改

#### 2.1 模型层修改
- **文件**: `server/model/smartcreate/u_game_user_auth_code.go`
- **修改**: 在RoleGameId字段的gorm标签中添加`uniqueIndex`

#### 2.2 服务层修改
- **文件**: `server/service/smartcreate/u_game_user_auth_code.go`
- **修改**: 
  - 添加fmt包导入
  - 在CreateGameUserAuthCode方法中添加role_game_id唯一性检查
  - 在UpdateGameUserAuthCode方法中添加role_game_id唯一性检查（排除当前记录）

#### 2.3 导入服务修改
- **文件**: `server/service/smartcreate/u_game_user_auth_code_import.go`
- **修改**: 在批量导入前添加role_game_id唯一性批量检查

## 执行步骤

### 步骤1: 检查重复数据
```sql
-- 先检查是否有重复的role_game_id
SELECT role_game_id, COUNT(*) as count
FROM game_user_auth_codes
WHERE role_game_id IS NOT NULL AND role_game_id != ''
GROUP BY role_game_id
HAVING COUNT(*) > 1;
```

### 步骤2: 处理重复数据（如有）
```sql
-- 删除重复数据，保留最早的一条
DELETE FROM game_user_auth_codes 
WHERE id NOT IN (
    SELECT MIN(id) 
    FROM game_user_auth_codes 
    WHERE role_game_id IS NOT NULL AND role_game_id != ''
    GROUP BY role_game_id
);
```

### 步骤3: 添加唯一索引
```sql
-- 添加唯一索引
CREATE UNIQUE INDEX idx_game_user_auth_codes_role_game_id ON game_user_auth_codes(role_game_id);
```

### 步骤4: 重启后端服务
由于修改了Go代码，需要重启后端服务以应用更改。

## 验证

### 创建测试
```bash
# 尝试创建重复的role_game_id
POST /game_user_auth_code/createGameUserAuthCode
{
  "roleGameId": "test123",
  ...其他字段
}

# 再次尝试相同的role_game_id，应该返回错误
```

### 导入测试
```bash
# 尝试导入包含重复role_game_id的Excel文件
# 应该返回具体的重复ID列表
```

## 错误处理

- **创建时重复**: 返回"角色游戏ID 'xxx' 已存在"
- **更新时重复**: 返回"角色游戏ID 'xxx' 已存在"
- **导入时重复**: 返回"以下角色游戏ID已存在: [xxx, yyy, zzz]"

## 注意事项

1. **数据备份**: 执行SQL修改前请备份数据库
2. **空值处理**: role_game_id为NULL或空字符串时不受唯一约束影响
3. **性能影响**: 唯一索引会略微影响写入性能，但大幅提升查询性能
4. **兼容性**: 前端无需修改，后端API接口保持不变