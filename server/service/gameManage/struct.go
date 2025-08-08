package gameManage

// Player 玩家数据结构体
// 用于存储从Server.html中提取的玩家信息

type Player struct {
	PlayerAccount            string `json:"player_account"`              // 玩家账号
	PlayerName               string `json:"player_name"`                 // 玩家名
	Attack                   string `json:"attack"`                      // 攻击力
	Level                    string `json:"level"`                       // 等级
	Recharge                 string `json:"recharge"`                    // 充值
	GoldCoin                 string `json:"gold_coin"`                   // 元宝
	GoldCoinAmount           string `json:"gold_coin_amount"`            // 元宝数量
	Online                   string `json:"online"`                      // 在线
	Muted                    string `json:"muted"`                       // 禁言
	Banned                   string `json:"banned"`                      // 封号
	TransactionCount         string `json:"transaction_count"`           // 交易次数
	TodayTransactionCount    string `json:"today_transaction_count"`     // 今天交易次数
	AuctionListCount         string `json:"auction_list_count"`          // 拍卖行上架次数
	TodayAuctionListCount    string `json:"today_auction_list_count"`    // 今天拍卖行上架次数
	AuctionSuccessCount      string `json:"auction_success_count"`       // 拍卖行竞拍成功次数
	TodayAuctionSuccessCount string `json:"today_auction_success_count"` // 今日拍卖行竞拍成功次数
	IsMicroClient            string `json:"is_micro_client"`             // 是否微端登录
	LastLoginTime            string `json:"last_login_time"`             // 最后登录时间
	LoginIP                  string `json:"login_ip"`                    // 登录ip
	LastLogoutTime           string `json:"last_logout_time"`            // 最后下线时间
}

// CSVPlayer CSV文件中的玩家详情数据结构体
// 用于存储从玩家详情CSV文件中提取的玩家信息
type CSVPlayer struct {
	PlayerAccount   string `json:"player_account"`   // 玩家账号
	PlayerName      string `json:"player_name"`      // 玩家名
	PlayerId        string `json:"player_id"`        // 玩家id
	PlayerIP        string `json:"player_ip"`        // 玩家ip
	Level           string `json:"level"`            // 等级
	LoginMap        string `json:"login_map"`        // 登录地图
	Recharge        string `json:"recharge"`         // 充值
	TransactionCount string `json:"transaction_count"` // 交易次数
	FirstServerId   string `json:"first_server_id"`  // 最初登录区服id
	OnlineStatus    string `json:"online_status"`    // 是否在线
	LastLoginTime   string `json:"last_login_time"`  // 最后登录时间
	LastLogoutTime  string `json:"last_logout_time"` // 最后登出时间
	CreateTime      string `json:"create_time"`      // 角色创建时间
}

// PlayersResponse 玩家数据响应结构体
// 用于存储多个玩家数据

type PlayersResponse struct {
	Players []Player `json:"players"` // 玩家列表
}

type NewServer struct {
	Id       string `json:"id"`
	ServerId string `json:"serverid"`
	Name     string `json:"name"`
}

// 如果是数组形式，需要外层包裹
type NewServerResponse struct {
	List []NewServer `json:"list"` // 假设外层有list字段包裹数组
	// 如果JSON直接是数组，则不需要这个结构体，直接解析到[]NewServer即可
}

// PlayerAction 玩家行为数据结构体
// 用于存储从CSV文件中提取的玩家行为信息
type PlayerAction struct {
	Time         string `json:"time"`           // 时间
	PlayerName   string `json:"player_name"`    // 角色名
	PlayerId     string `json:"player_id"`      // 角色id
	Level        string `json:"level"`          // 等级
	Map          string `json:"map"`            // 地图
	Action       string `json:"action"`         // 行为
	ItemId       string `json:"item_id"`        // 物品id
	Item         string `json:"item"`           // 物品
	BeforeCount  string `json:"before_count"`   // 变化前数量
	ChangeCount  string `json:"change_count"`   // 变化数量
	CurrentCount string `json:"current_count"`  // 当前数量
	Related      string `json:"related"`        // Related
	ItemUniqueId string `json:"item_unique_id"` // 物品唯一ID
}

// ServerInfo 存储服务器信息结构体
type ServerInfo struct {
	ZoneName    string `json:"zone_name"`    // 区名
	Zone        string `json:"zone"`         // 区
	MergedZones string `json:"merged_zones"` // 被合区
	OpenTime    string `json:"open_time"`    // 开服时间
}
