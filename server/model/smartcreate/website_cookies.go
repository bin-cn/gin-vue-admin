// 自动生成模板CookieData
package smartcreate

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 网站cookie 结构体  CookieData
type CookieData struct {
	global.GVA_MODEL
	BusinessCode *int       `json:"businessCode" form:"businessCode" gorm:"comment:业务类型标识;column:business_code;'search';"`            //业务编号
	CookieValue  *string    `json:"cookieValue" form:"cookieValue" gorm:"comment:cookie的具体值;column:cookie_value;" binding:"required"` //cookie值
	Status       *int       `json:"status" form:"status" gorm:"comment:cookie状态，如1为正常，0为失效;column:status;"`                           //状态
	Remark       *string    `json:"remark" form:"remark" gorm:"comment:备注信息;column:remark;"`                                          //备注信息
	LastUseTime  *time.Time `json:"lastUseTime" form:"lastUseTime" gorm:"comment:最后一次使用时间;column:last_use_time;"`                     //最后一次使用时间
}

// TableName 网站cookie CookieData自定义表名 website_cookies
func (CookieData) TableName() string {
	return "website_cookies"
}
