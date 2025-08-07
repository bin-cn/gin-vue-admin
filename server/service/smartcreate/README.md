# 用户授权码自定义Excel导入功能

## 功能概述

这个自定义导入功能专门为`templateID=smartcreate_GameUserAuthCode`设计，解决了以下两个问题：

1. **登录码唯一性处理**：当导入的Excel中存在重复的`login_code`时，系统会自动生成新的唯一登录码
2. **用户验证**：根据`assigner_name`字段自动查询`sys_users`表中的`nick_name`，验证用户是否存在并获取对应的`user_id`

## 使用方式

### 1. 下载导入模板

```bash
GET /smartcreate/gameUserAuthCode/downloadTemplate
```

### 2. 自定义导入

```bash
POST /smartcreate/gameUserAuthCode/importCustom
Content-Type: multipart/form-data

参数：
- file: Excel文件（.xlsx或.xls格式）
```

### 3. Excel模板格式

| 列名 | 说明 | 是否必填 |
|------|------|----------|
| login_code | 登录码 | 是 |
| assigner_name | 使用人（对应sys_users.nick_name） | 是 |
| machine_no_name | 机器编号 | 否 |
| game_server_name | 游戏服务器名称 | 否 |
| account | 账号 | 否 |
| role_game_id | 游戏角色ID | 否 |
| i_d_name | 姓名 | 否 |
| i_d_card_number | 身份证号 | 否 |

## 处理逻辑

### 登录码唯一性处理
- 如果导入的`login_code`已存在，系统会自动在末尾添加`_随机数`后缀
- 例如：`TEST001` → `TEST001_1234`

### 用户验证流程
1. 根据`assigner_name`查询`sys_users`表的`nick_name`字段
2. 如果用户不存在，返回错误信息："用户不存在: [assigner_name]"
3. 如果用户存在，获取对应的`id`作为`user_id`

## 错误处理

系统会返回详细的错误信息，包括：
- 哪一行数据出现问题
- 具体的错误原因
- 缺失的必填字段
- 用户不存在的信息

## 安全特性

- **完全隔离**：所有代码都在`smartcreate`模块内，不影响原有系统
- **事务处理**：使用数据库事务，确保数据一致性
- **错误回滚**：任何错误都会回滚整个导入操作
- **日志记录**：详细记录导入过程中的所有操作

## 示例

### 成功导入示例
```json
{
  "code": 0,
  "msg": "导入成功",
  "data": {}
}
```

### 错误示例
```json
{
  "code": 7,
  "msg": "第3行: 用户不存在: 张三",
  "data": {}
}
```

## 注意事项

1. 确保Excel文件格式正确
2. 所有必填字段必须提供
3. 使用人名称必须与`sys_users.nick_name`完全匹配
4. 建议先下载模板文件，按格式填写后再导入
5. 大文件导入可能需要较长时间，请耐心等待

## 技术实现

- 服务层：`GameUserAuthCodeImportService`
- API层：`GameUserAuthCodeImportApi`
- 路由：`/smartcreate/gameUserAuthCode/`
- 数据库表：`game_user_auth_code`和`sys_users`