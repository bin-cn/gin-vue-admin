package gameManage

import (
	"fmt"
	"net/url"
	"time"
)

// BuildTime 根据输入的分钟数生成符合要求的时间范围
// minutes: 实际要查询的分钟数，必须是5的倍数
// 返回值: 格式化后的开始时间和结束时间字符串（格式：2025-08-04+20:00）
func BuildTime(minutes int) (string, string) {
	// 获取当前时间
	now := time.Now()

	// 确保分钟数是5的倍数
	if minutes%5 != 0 {
		minutes = (minutes/5 + 1) * 5
	}

	// 计算结束时间: 向下取整到最近的5分钟倍数
	minute := now.Minute()
	roundedMinute := (minute / 5) * 5
	endTime := time.Date(
		now.Year(), now.Month(), now.Day(),
		now.Hour(), roundedMinute, 0, 0,
		now.Location(),
	)

	// 计算开始时间: 结束时间减去指定的分钟数
	duration := time.Duration(minutes) * time.Minute
	startTime := endTime.Add(-duration)

	// 确保开始时间不超过当前时间
	if startTime.After(now) {
		startTime = now.Add(-duration)
	}

	// 手动构建时间字符串，格式为2025-08-07+17:40，让URL编码自动处理冒号为%3A
	startStr := fmt.Sprintf(
		"%04d-%02d-%02d %02d:%02d",
		startTime.Year(), startTime.Month(), startTime.Day(),
		startTime.Hour(), startTime.Minute(),
	)
	endStr := fmt.Sprintf(
		"%04d-%02d-%02d %02d:%02d",
		endTime.Year(), endTime.Month(), endTime.Day(),
		endTime.Hour(), endTime.Minute(),
	)

	return startStr, endStr
}

// FormatItem 格式化物品参数
// item: 原始物品字符串
// 返回值: 格式化并URL编码后的物品字符串
func FormatItem(item string) string {
	// 如果item已经是URL编码格式，则直接返回
	// 否则，对其进行编码
	if item != "" {
		// 检查是否已经编码
		decoded, err := url.QueryUnescape(item)
		if err != nil || decoded == item {
			// 未编码或解码失败，进行编码
			encoded := url.QueryEscape(item)
			return encoded
		}
	}
	return item
}
