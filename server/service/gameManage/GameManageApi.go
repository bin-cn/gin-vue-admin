package gameManage

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"github.com/flipped-aurora/gin-vue-admin/server/service/gameManage/request"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// GameManager 游戏管理客户端结构体，用于封装相关请求方法和cookie
type GameManager struct {
	cookie string // 存储当前使用的cookie
}

// NewGameManager 创建一个新的GameManager实例
// 可传入初始cookie，若为空则使用默认cookie
func NewGameManager(initialCookie string) (*GameManager, error) {
	// 检查cookie是否为空
	if strings.TrimSpace(initialCookie) == "" {
		return nil, errors.New("必须提供有效的cookie，不能为空")
	}
	// 返回初始化后的实例
	return &GameManager{
		cookie: initialCookie,
	}, nil
}

// SetCookie 设置当前实例使用的cookie
func (gm *GameManager) SetCookie(cookie string) {
	gm.cookie = cookie
}

// Cookie 获取当前实例使用的cookie
func (gm *GameManager) Cookie() string {
	return gm.cookie
}

// baseRequestMSCS 基础请求方法，封装HTTP请求逻辑
func (gm *GameManager) baseRequestMSCS(urlStr string, method string, headers map[string]string, params map[string]string) (string, error) {
	var result string = ""

	// 根据请求方法决定数据格式和请求体数据
	var format request.DataFormat
	var data any = nil

	if method == "POST" {
		// 对于POST请求，根据Content-Type决定格式
		if headers["Content-Type"] == "application/x-www-form-urlencoded; charset=UTF-8" {
			format = request.FormFormat
		} else {
			format = request.JSONFormat
		}
		data = params
	} else {
		// 对于非POST请求（如GET），使用表单格式，且请求体数据为nil
		format = request.FormFormat
		data = nil

	}

	// 调用request包的HttpRequest方法发送请求
	resp, err := request.HttpRequest(urlStr, method, headers, params, data, format)

	if err != nil {
		return "", fmt.Errorf("请求发送失败: %w", err)
	}
	defer resp.Body.Close()

	// 处理响应内容（gzip解压和编码转换）
	var reader io.Reader = resp.Body
	contentEncoding := resp.Header.Get("Content-Encoding")
	if strings.Contains(contentEncoding, "gzip") {
		gzReader, err := gzip.NewReader(reader)
		if err != nil {
			return result, fmt.Errorf("gzip解压失败: %w", err)
		}
		defer gzReader.Close()
		reader = gzReader
	}

	// 处理application/octet-stream类型的响应，特别是GBK编码
	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/octet-stream") {
		// 读取原始字节
		bodyBytes, err := io.ReadAll(reader)
		if err != nil {
			return result, fmt.Errorf("读取响应体失败: %w", err)
		}

		// 检查是否包含GBK编码标识
		if strings.Contains(contentType, "charset=gbk") {
			// 尝试使用GBK编码解码
			decoder := simplifiedchinese.GBK.NewDecoder()
			decodedBytes, err := io.ReadAll(decoder.Reader(bytes.NewReader(bodyBytes)))
			if err != nil {
				// 如果解码失败，使用原始字节
				result = string(bodyBytes)
			} else {
				result = string(decodedBytes)
			}
		} else {
			// 假设是UTF-8编码
			result = string(bodyBytes)
		}

		return result, nil
	}

	// 对于其他类型的响应，保持当前reader不变

	// 读取响应内容
	body, err := io.ReadAll(reader)
	if err != nil {
		return result, fmt.Errorf("读取响应体失败: %w", err)
	}

	strBody := string(body)
	return strBody, nil
}

// baseHeader 构建基础请求头
func (gm *GameManager) baseHeader() map[string]string {
	return map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9",
		"Connection":       "keep-alive",
		"Cookie":           gm.cookie, // 使用结构体中存储的cookie
		"Host":             "ht.sqdk.yuhetx.net",
		"Referer":          "http://ht.sqdk.yuhetx.net/player/oneband",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}
}

// GetSelectServers 根据区服ID,获取服务器内部ID列表,玩家详情，玩家ID中使用的服务器ID
func (gm *GameManager) GetSelectServers() (string, error) {
	urlStr := "http://ht.sqdk.yuhetx.net/default/GetSelectServers"
	method := "GET"
	header := gm.baseHeader()
	params := map[string]string{
		"sSearch": "",
		"page":    "1",
		"oper":    "ad_93772",
		// "_":       "1754641339896",
	}
	header["Referer"] = "http://ht.sqdk.yuhetx.net/player/onelist"

	res, err := gm.baseRequestMSCS(urlStr, method, header, params)
	if err != nil {
		fmt.Printf("获取服务器列表失败: %v\n", err)
		return res, err
	}
	return res, nil
}

// Operservers 获取运营服务器列表--->获取区服ID对应的服务器ID.
// 返回值：
//
//	[]NewServer：服务器列表（包含每个服务器的ID、ServerId和名称）
//	error：请求或解析失败时返回的错误信息（成功时为nil）
func (gm *GameManager) Operservers() ([]NewServer, error) {
	urlStr := "http://ht.sqdk.yuhetx.net/helper/operservers"
	method := "GET"
	header := gm.baseHeader()
	params := map[string]string{
		"oper": "ad_93772",
	}

	var servers []NewServer

	res, err := gm.baseRequestMSCS(urlStr, method, header, params)
	if err != nil {
		fmt.Printf("获取服务器列表失败: %v\n", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &servers)
	if err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return servers, err
	}
	return servers, nil
}

// Login 处理登录和获取验证码功能
// username: 用户名
// password: 密码
// code: 验证码(登录时需要，获取验证码时可空)
// operate: 操作类型("getcode"获取验证码, "checkuser"登录)
// 返回值: 响应结果和可能的错误
func (gm *GameManager) login(username, password, code, operate string) (string, error) {
	urlStr := "http://ht.sqdk.yuhetx.net/default/login"
	method := "POST"
	header := gm.baseHeader()
	// 设置Content-Type为表单格式
	header["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	// 更新Referer
	header["Referer"] = "http://ht.sqdk.yuhetx.net/default/login"

	params := map[string]string{
		"name":    username,
		"passwd":  password,
		"operate": operate,
	}

	// 如果提供了验证码，则添加到参数中
	if code != "" {
		params["code"] = code
	}

	res, err := gm.baseRequestMSCS(urlStr, method, header, params)
	if err != nil {
		fmt.Printf("%s失败: %v\n", map[string]string{"getcode": "获取验证码", "checkuser": "登录"}[operate], err)
		return "", err
	}
	return res, nil
}

func (gm *GameManager) SendCode(username, password string) (bool, error) {
	var sendOK bool = false
	res, err := gm.login(username, password, "", "getcode")
	if err != nil {
		fmt.Printf("SendCode 失败: %v,%s\n", err, res)
		return sendOK, err
	}
	// 判断是否包含验证码发送成功的提示
	if strings.Contains(res, "验证码已发送至您的手机，请注意查收") {
		print("验证码发送成功")
		sendOK = true
	}

	return sendOK, err
}

func (gm *GameManager) Login_enent(username, password, code string) (bool, error) {
	var sendOK bool = false
	res, err := gm.login(username, password, code, "checkuser")

	if err != nil {
		fmt.Printf("SendCode 失败: %v,%s\n", err, res)
		return sendOK, err
	}

	// 判断是否包含验证码发送成功的提示
	if strings.Contains(res, "success") {
		print("登录成功")
		sendOK = true
	}

	return sendOK, err
}

// 玩家详情(封禁)-选区并返回查询的所有数据
func (gm *GameManager) Server(sid string, isMaster bool) ([]Player, error) {
	urlStr := "http://ht.sqdk.yuhetx.net/default/server"
	method := "GET"
	header := gm.baseHeader()
	params := map[string]string{
		"oper": "ad_93772",
		"sid":  sid,
	}
	header["Referer"] = "http://ht.sqdk.yuhetx.net/player/oneband?rolename=&loginname=&action=oneband"

	// 是否切主区服
	if isMaster {
		params["getMainServer"] = "getMainServer"
	}

	res, err := gm.baseRequestMSCS(urlStr, method, header, params)
	if err != nil {
		fmt.Printf("获取服务器列表失败: %v,%v\n", err, res)
		return nil, err
	}

	var players []Player
	players, err = parsePlayerHTML(res)
	if err != nil {
		fmt.Printf("解析HTML 数据失败.请联系管理员:  %v,%v\n", err, res)
		return nil, err
	}
	return players, nil
}

// 辅助函数：返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 玩家详情-选区并返回查询的所有数据
func (gm *GameManager) OneList(sid string, isMaster bool) ([]CSVPlayer, error) {
	urlStr := "http://ht.sqdk.yuhetx.net/player/onelist"
	method := "GET"
	header := gm.baseHeader()
	params := map[string]string{
		"action":    "onelist",
		"rolename":  "",
		"loginname": "",
		"export":    "1",
		"sid":       sid,
	}
	header["Referer"] = "http://ht.sqdk.yuhetx.net/player/onelist?rolename=&loginname=&action=onelist"

	res, err := gm.baseRequestMSCS(urlStr, method, header, params)
	if err != nil {
		fmt.Printf("获取玩家列表失败: %v,%v\n", err, res)
		return nil, err
	}

	// 解析返回的CSV数据
	players, err := parsePlayerCSV(res)
	if err != nil {
		return nil, fmt.Errorf("解析CSV数据失败: %w", err)
	}
	return players, nil
}

// GetPlayerActions 获取玩家行为数据
// rolename: 角色名
// roleid: 角色ID
// starttime: 开始时间(格式: 2025-08-04+20%3A00)，为空则使用默认时间
// endtime: 结束时间(格式: 2025-08-04+20%3A35)，为空则使用默认时间
// item: 物品(格式: 20%3A%E5%85%83%E5%AE%9D)，会自动进行URL编码
// 返回值: 玩家行为数据数组和可能的错误
func (gm *GameManager) GetPlayerActions(starttime, endtime, item string) ([]PlayerAction, error) {
	urlStr := "http://ht.sqdk.yuhetx.net/player/action"
	method := "GET"

	// 处理时间参数，如果为空则使用默认时间范围(最近10分钟)
	if starttime == "" || endtime == "" {
		starttime, endtime = BuildTime(1) // 默认使用1，表示最近10分钟
	}

	// 格式化物品参数
	// item = FormatItem(item)

	// 创建自定义请求头
	header := map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Encoding":           "gzip, deflate",
		"Accept-Language":           "zh-CN,zh;q=0.9",
		"Connection":                "keep-alive",
		"Cookie":                    gm.cookie,
		"Host":                      "ht.sqdk.yuhetx.net",
		"Referer":                   "http://ht.sqdk.yuhetx.net/player/action",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36",
	}

	// 设置请求参数
	params := map[string]string{
		"rolename":  "",
		"roleid":    "",
		"starttime": starttime,
		"endtime":   endtime,
		"item":      item,
		"type":      "",
		"makeid":    "",
		"action":    "searchAction",
		"export":    "1",
	}

	// 发送请求
	res, err := gm.baseRequestMSCS(urlStr, method, header, params)
	if err != nil {
		fmt.Printf("获取玩家行为数据失败: %v\n", err)
		return nil, err
	}

	// 解析CSV数据
	actions, err := parsePlayerActionCSV(res)
	if err != nil {
		fmt.Printf("解析CSV数据失败: %v\n", err)
		return nil, err
	}

	return actions, nil
}

// GetServerList 获取服务器----合区信息列表
func (gm *GameManager) GetServerList() ([]ServerInfo, error) {
	url := "http://ht.sqdk.yuhetx.net/manager/servers?search=ad_93772"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	headers := map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Encoding":           "gzip, deflate",
		"Accept-Language":           "zh-CN,zh;q=0.9",
		"Connection":                "keep-alive",
		"Host":                      "ht.sqdk.yuhetx.net",
		"Referer":                   "http://ht.sqdk.yuhetx.net/manager/servers",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36",
		"Cookie":                    gm.cookie,
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	// req.Header.Add("Cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer res.Body.Close()

	// 处理gzip压缩
	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return nil, fmt.Errorf("创建gzip读取器失败: %w", err)
		}
		defer reader.Close()
	default:
		reader = res.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 网页已经是UTF-8编码，无需转换
	htmlStr := string(body)
	return parseServerListHTML(htmlStr)
}

// ParsePlayerCSV 解析CSV格式的玩家详情数据
// csvContent: CSV格式的玩家详情数据字符串
// 返回值: CSV玩家详情数据数组和可能的错误
func parsePlayerCSV(csvContent string) ([]CSVPlayer, error) {
	var players []CSVPlayer

	// 使用标准库encoding/csv来正确解析CSV
	reader := csv.NewReader(strings.NewReader(csvContent))

	// 设置CSV解析选项
	reader.FieldsPerRecord = -1    // 允许可变字段数量
	reader.LazyQuotes = true       // 允许不严格的引号
	reader.TrimLeadingSpace = true // 去除前导空格

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("解析CSV数据失败: %w", err)
	}

	// 跳过标题行，从第二行开始
	if len(records) <= 1 {
		return players, nil
	}

	// 解析每一行数据
	for _, record := range records[1:] {
		// 跳过空行或只有空白字符的行
		if len(record) == 0 || (len(record) == 1 && strings.TrimSpace(record[0]) == "") {
			continue
		}

		// 创建CSVPlayer实例，根据实际字段数量处理
		player := CSVPlayer{}

		// 安全地填充字段，避免索引越界
		if len(record) > 0 {
			player.PlayerAccount = strings.TrimSpace(record[0])
		}
		if len(record) > 1 {
			player.PlayerName = strings.TrimSpace(record[1])
		}
		if len(record) > 2 {
			player.PlayerId = strings.TrimSpace(record[2])
		}
		if len(record) > 3 {
			player.PlayerIP = strings.TrimSpace(record[3])
		}
		if len(record) > 4 {
			player.Level = strings.TrimSpace(record[4])
		}
		if len(record) > 5 {
			player.LoginMap = strings.TrimSpace(record[5])
		}
		if len(record) > 6 {
			player.Recharge = strings.TrimSpace(record[6])
		}
		if len(record) > 7 {
			player.TransactionCount = strings.TrimSpace(record[7])
		}
		if len(record) > 8 {
			player.FirstServerId = strings.TrimSpace(record[8])
		}
		if len(record) > 9 {
			player.OnlineStatus = strings.TrimSpace(record[9])
		}
		if len(record) > 10 {
			player.LastLoginTime = strings.TrimSpace(record[10])
		}
		if len(record) > 11 {
			player.LastLogoutTime = strings.TrimSpace(record[11])
		}
		if len(record) > 12 {
			player.CreateTime = strings.TrimSpace(record[12])
		}

		// 添加到结果数组
		players = append(players, player)
	}

	return players, nil
}

// ParsePlayerActionCSV 解析CSV格式的玩家行为数据
// csvContent: CSV格式的玩家行为数据字符串
// 返回值: 玩家行为数据数组和可能的错误
func parsePlayerActionCSV(csvContent string) ([]PlayerAction, error) {
	var actions []PlayerAction

	// 使用标准库encoding/csv来正确解析CSV
	reader := csv.NewReader(strings.NewReader(csvContent))

	// 设置CSV解析选项
	reader.FieldsPerRecord = -1    // 允许可变字段数量
	reader.LazyQuotes = true       // 允许不严格的引号
	reader.TrimLeadingSpace = true // 去除前导空格

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("解析CSV数据失败: %w", err)
	}

	// 跳过标题行，从第二行开始
	if len(records) <= 1 {
		return actions, nil
	}

	// 解析每一行数据
	for _, record := range records[1:] {
		// 跳过空行或只有空白字符的行
		if len(record) == 0 || (len(record) == 1 && strings.TrimSpace(record[0]) == "") {
			continue
		}

		// 创建PlayerAction实例，根据实际字段数量处理
		action := PlayerAction{}

		// 安全地填充字段，避免索引越界
		if len(record) > 0 {
			action.Time = strings.TrimSpace(record[0])
		}
		if len(record) > 1 {
			action.PlayerName = strings.TrimSpace(record[1])
		}
		if len(record) > 2 {
			action.PlayerId = strings.TrimSpace(record[2])
		}
		if len(record) > 3 {
			action.Level = strings.TrimSpace(record[3])
		}
		if len(record) > 4 {
			action.Map = strings.TrimSpace(record[4])
		}
		if len(record) > 5 {
			action.Action = strings.TrimSpace(record[5])
		}
		if len(record) > 6 {
			action.ItemId = strings.TrimSpace(record[6])
		}
		if len(record) > 7 {
			action.Item = strings.TrimSpace(record[7])
		}
		if len(record) > 8 {
			action.BeforeCount = strings.TrimSpace(record[8])
		}
		if len(record) > 9 {
			action.ChangeCount = strings.TrimSpace(record[9])
		}
		if len(record) > 10 {
			action.CurrentCount = strings.TrimSpace(record[10])
		}
		if len(record) > 11 {
			action.Related = strings.TrimSpace(record[11])
		}
		if len(record) > 12 {
			action.ItemUniqueId = strings.TrimSpace(record[12])
		}

		// 添加到结果数组
		actions = append(actions, action)
	}

	return actions, nil
}
