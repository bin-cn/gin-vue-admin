package service

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"time"
)

type ClearTableTask struct {
}

// @description: 清理数据库表数据
// tableName(表名) string, compareField(比较字段) string, interval(间隔) string
// @return: error
func (t ClearTableTask) Exec(arg interface{}) error {
	var ClearTableDetail []common.ClearDB

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "sys_operation_records",
		CompareField: "created_at",
		Interval:     "2160h",
	})

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "jwt_blacklists",
		CompareField: "created_at",
		Interval:     "168h",
	})

	if global.GVA_DB == nil {
		return errors.New("db Cannot be empty")
	}

	for _, detail := range ClearTableDetail {
		duration, err := time.ParseDuration(detail.Interval)
		if err != nil {
			return err
		}
		if duration < 0 {
			return errors.New("parse duration < 0")
		}
		err = global.GVA_DB.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", detail.TableName, detail.CompareField), time.Now().Add(-duration)).Error
		if err != nil {
			return err
		}
	}
	return nil
}
