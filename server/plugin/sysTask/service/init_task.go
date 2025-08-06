package service

var JobList map[string]JobExec

// InitJob
// 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
func InitJob() {
	JobList = map[string]JobExec{
		"ExamplesOne":    ExamplesOne{},
		"ClearTableTask": ClearTableTask{},
		// ...
	}
}
