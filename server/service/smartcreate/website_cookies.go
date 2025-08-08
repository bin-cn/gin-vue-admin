
package smartcreate

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate"
    smartcreateReq "github.com/flipped-aurora/gin-vue-admin/server/model/smartcreate/request"
)

type CookieDataService struct {}
// CreateCookieData 创建网站cookie记录
// Author [yourname](https://github.com/yourname)
func (wcService *CookieDataService) CreateCookieData(ctx context.Context, wc *smartcreate.CookieData) (err error) {
	err = global.GVA_DB.Create(wc).Error
	return err
}

// DeleteCookieData 删除网站cookie记录
// Author [yourname](https://github.com/yourname)
func (wcService *CookieDataService)DeleteCookieData(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&smartcreate.CookieData{},"id = ?",ID).Error
	return err
}

// DeleteCookieDataByIds 批量删除网站cookie记录
// Author [yourname](https://github.com/yourname)
func (wcService *CookieDataService)DeleteCookieDataByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]smartcreate.CookieData{},"id in ?",IDs).Error
	return err
}

// UpdateCookieData 更新网站cookie记录
// Author [yourname](https://github.com/yourname)
func (wcService *CookieDataService)UpdateCookieData(ctx context.Context, wc smartcreate.CookieData) (err error) {
	err = global.GVA_DB.Model(&smartcreate.CookieData{}).Where("id = ?",wc.ID).Updates(&wc).Error
	return err
}

// GetCookieData 根据ID获取网站cookie记录
// Author [yourname](https://github.com/yourname)
func (wcService *CookieDataService)GetCookieData(ctx context.Context, ID string) (wc smartcreate.CookieData, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wc).Error
	return
}
// GetCookieDataInfoList 分页获取网站cookie记录
// Author [yourname](https://github.com/yourname)
func (wcService *CookieDataService)GetCookieDataInfoList(ctx context.Context, info smartcreateReq.CookieDataSearch) (list []smartcreate.CookieData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&smartcreate.CookieData{})
    var wcs []smartcreate.CookieData
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["id"] = true
           orderMap["created_at"] = true
         	orderMap["business_code"] = true
         	orderMap["last_use_time"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wcs).Error
	return  wcs, total, err
}
func (wcService *CookieDataService)GetCookieDataPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
