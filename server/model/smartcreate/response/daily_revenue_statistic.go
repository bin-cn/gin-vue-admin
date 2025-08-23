package response

type DailyRevenueStatistic struct {
	StatisticDate *string `json:"statisticDate" ` //统计日期
	UserId        *string `json:"userId" `        //用户ID
	UserNickname  *string `json:"userNickname" `  //用户昵称
	ServerName    *string `json:"serverName" `    //区服名字
	ServerZoneId  *string `json:"serverZoneId" `  //区服ID
	StatisticType *string `json:"statisticType" ` //统计类型
	Amount        *int64  `json:"amount" `        //产出金额
}
