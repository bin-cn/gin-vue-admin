package gameManage

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParsePlayerHTML 解析Server.html中的玩家数据
// htmlContent: Server.html的HTML内容
// 返回: 玩家数据数组和可能的错误
func parsePlayerHTML(htmlContent string) ([]Player, error) {
	// 检查输入是否为空
	if htmlContent == "" {
		return nil, fmt.Errorf("HTML内容不能为空")
	}

	// 创建一个字符串读取器
	reader := strings.NewReader(htmlContent)

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 存储玩家数据的切片
	var players []Player

	// 查找表格并提取数据
	doc.Find("table#sample_editable_1 tbody tr").Each(func(i int, s *goquery.Selection) {
		// 创建一个新的Player实例
		player := Player{}

		// 提取表格单元格数据
		cells := s.Find("td")

		// 根据单元格索引设置Player字段
		// 注意: 这里的索引需要与表格中的列顺序匹配
		if cells.Length() >= 20 {
			player.PlayerAccount = strings.TrimSpace(cells.Eq(0).Text())
			player.PlayerName = strings.TrimSpace(cells.Eq(1).Text())
			player.Attack = strings.TrimSpace(cells.Eq(2).Text())
			player.Level = strings.TrimSpace(cells.Eq(3).Text())
			player.Recharge = strings.TrimSpace(cells.Eq(4).Text())
			player.GoldCoin = strings.TrimSpace(cells.Eq(5).Text())
			player.GoldCoinAmount = strings.TrimSpace(cells.Eq(6).Text())
			player.Online = strings.TrimSpace(cells.Eq(7).Text())
			player.Muted = strings.TrimSpace(cells.Eq(8).Text())
			player.Banned = strings.TrimSpace(cells.Eq(9).Text())
			player.TransactionCount = strings.TrimSpace(cells.Eq(10).Text())
			player.TodayTransactionCount = strings.TrimSpace(cells.Eq(11).Text())
			player.AuctionListCount = strings.TrimSpace(cells.Eq(12).Text())
			player.TodayAuctionListCount = strings.TrimSpace(cells.Eq(13).Text())
			player.AuctionSuccessCount = strings.TrimSpace(cells.Eq(14).Text())
			player.TodayAuctionSuccessCount = strings.TrimSpace(cells.Eq(15).Text())
			player.IsMicroClient = strings.TrimSpace(cells.Eq(16).Text())
			player.LastLoginTime = strings.TrimSpace(cells.Eq(17).Text())
			player.LoginIP = strings.TrimSpace(cells.Eq(18).Text())
			player.LastLogoutTime = strings.TrimSpace(cells.Eq(19).Text())
			// 添加到玩家切片
			players = append(players, player)
		}
	})

	// 检查是否提取到数据
	if len(players) == 0 {
		return nil, fmt.Errorf("未找到玩家数据")
	}

	return players, nil
}

// GetPlayersFromHTMLFile 从HTML文件中获取玩家数据
// filePath: HTML文件路径
// 返回: 玩家数据响应和可能的错误
func getPlayersFromHTMLFile(filePath string) ([]Player, error) {
	// 读取HTML文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("读取HTML文件失败: %w", err)
	}

	// 调用ParsePlayerHTML解析HTML内容
	players, err := parsePlayerHTML(string(content))
	if err != nil {
		return nil, err
	}

	// 返回玩家数据响应
	return players, nil
}
