package response

// GameUserStatsResp 用户元宝统计响应对象
// 方便后续扩展更多统计字段
type GameUserStatsResp struct {
	OnlineIngotTotal    int64 `json:"onlineIngotTotal"`    // 在线元宝总数
	UnBoundIngotQuantity int64 `json:"unBoundIngotQuantity"` // 未绑定元宝数量
}