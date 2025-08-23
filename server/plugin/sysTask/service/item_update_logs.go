package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service/smartcreate"
)

type ItemUpdateLogs struct{}

// @description: 物品更新日志
// @param: arg interface{}
// @return: error
func (t ItemUpdateLogs) Exec(arg interface{}) error {
	//arg 是字符串用逗号分隔的,需要解析为数组对象

	argStr := arg.(string)
	argArr := strings.Split(argStr, "|")
	if len(argArr) != 2 {
		return fmt.Errorf("物品更新日志任务参数错误,参数格式: 任务ID,角色ID")
	}
	cookieID := argArr[0]
	buildTileId := argArr[1]

	gm, err := smartcreate.VerifyItemUpdateLogsData(cookieID)
	if err != nil {
		return fmt.Errorf("物品更新日志任务验证数据失败: %v", err)
	}
	// 1.查询game_user_auth_code表中的所有main_server_zone字段值
	var mainServerZones []string
	result := global.GVA_DB.Raw("SELECT main_server_zone FROM gva.game_user_auth_code group by main_server_zone order by main_server_zone asc;").Scan(&mainServerZones)
	if result.Error != nil {
		return fmt.Errorf("查询game_user_auth_code表中的所有main_server_zone字段值失败: %v", result.Error)
	}

	// 转换一下类型
	minutes, err := strconv.ParseInt(buildTileId, 10, 32)
	if err != nil {
		return fmt.Errorf("buildTileId转换失败: %v", err)
	}

	//2. mainServerZones    这里拿到的是区服ID列表，需要从game_server表中查询对应的主区服ID,需要去重
	var mainServerZoneIds []int
	result = global.GVA_DB.Raw("SELECT DISTINCT main_server_id FROM gva.game_server WHERE main_server_zone_id IN ?", mainServerZones).Scan(&mainServerZoneIds)
	if result.Error != nil {
		return fmt.Errorf("查询game_server表中的所有main_server_zone字段值失败: %v", result.Error)
	}

	var list_errinfo []string

	// 3.遍历mainServerZones，调用gm.GetPlayerActions()方法
	for _, mainServerZone := range mainServerZoneIds {
		mainServerZoneStr := strconv.Itoa(mainServerZone)
		err := smartcreate.Update_player_info_by_serverid_lock(gm, smartcreate.ActionTypeScheduled, mainServerZoneStr, int(minutes))
		if err != nil {
			var myErr *smartcreate.MyError
			if errors.As(err, &myErr) {
				msginfo := fmt.Sprintf("区服[%s],更新失败:原因是 %s", mainServerZoneStr, myErr.Msg)
				list_errinfo = append(list_errinfo, msginfo)
				fmt.Printf("该区服[%s]更新失败,错误码:%d,错误信息:%s", mainServerZoneStr, myErr.Code, myErr.Msg)
				continue
			} else {
				return fmt.Errorf("更新区服[%s]失败: %v", mainServerZoneStr, err)
			}
		}

	}

	if len(list_errinfo) > 0 {
		return fmt.Errorf("更新完毕,以下区服未能更新成功:%s", strings.Join(list_errinfo, ","))
	}

	drService := &smartcreate.DailyRevenueRecordService{}
	err = drService.InstallDailyRevenueRecord()
	if err != nil {
		return fmt.Errorf("更新完毕.但是插入统计数据失败:%s", err.Error())
	}

	return nil
}
