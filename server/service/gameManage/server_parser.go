package gameManage

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParseServerListHTML 解析服务器列表HTML并返回服务器信息数组
func parseServerListHTML(htmlContent string) ([]ServerInfo, error) {
	var servers []ServerInfo

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找表格中的每一行数据（跳过表头）
	doc.Find("table#sample_editable_1 tbody tr").Each(func(i int, s *goquery.Selection) {
		var server ServerInfo

		// 提取每个单元格的数据
		tds := s.Find("td")

		if tds.Length() >= 4 {
			// 区名（第1列）
			server.ZoneName = strings.TrimSpace(tds.Eq(0).Text())

			// 区（第2列）
			server.Zone = strings.TrimSpace(tds.Eq(1).Text())

			// 被合区（第3列）
			server.MergedZones = strings.TrimSpace(tds.Eq(2).Text())

			// 开服时间（第4列）
			server.OpenTime = strings.TrimSpace(tds.Eq(3).Text())

			// 清理空白字符和特殊字符
			server.ZoneName = cleanText(server.ZoneName)
			server.Zone = cleanText(server.Zone)
			server.MergedZones = cleanText(server.MergedZones)
			server.OpenTime = cleanText(server.OpenTime)

			// 跳过空行
			if server.ZoneName != "" && server.Zone != "" {
				servers = append(servers, server)
			}
		}
	})

	return servers, nil
}

// cleanText 清理文本中的空白字符和特殊字符
func cleanText(text string) string {
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.TrimSpace(text)
	return text
}
