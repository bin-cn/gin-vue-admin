package service

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExamplesOne
// 新添加的job 必须按照以下格式定义，并实现Exec函数
type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg interface{}) error {
	global.GVA_LOG.Info("JobCore ExamplesOne exec success")
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	switch arg.(type) {

	case string:
		if arg.(string) != "" {
			fmt.Println("string", arg.(string))
		} else {
			fmt.Println("arg is nil")
		}
		break
	}

	return nil
}
