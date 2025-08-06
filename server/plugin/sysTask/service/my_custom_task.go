package service

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// MyCustomTask 自定义任务
type MyCustomTask struct{}

func (t MyCustomTask) Exec(arg interface{}) error {
	global.GVA_LOG.Info("MyCustomTask exec success")

	// 根据参数类型处理参数
	switch arg.(type) {
	case string:
		param := arg.(string)
		if param != "" {
			// 在这里编写您的业务逻辑
			global.GVA_LOG.Info("任务参数: " + param)
		}
	}
	print("任务执行成功", arg)
	return nil
}
