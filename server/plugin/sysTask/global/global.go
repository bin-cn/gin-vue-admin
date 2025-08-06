package global

const HttpExecTimeout = 300 // http任务执行时间不超过300秒
type SysTaskStatus int8

const (
	Disabled SysTaskStatus = 99 // 禁用
	Failure  SysTaskStatus = 0  // 失败
	Enabled  SysTaskStatus = 1  // 启用
	Running  SysTaskStatus = 10 // 运行中
	Finish   SysTaskStatus = 2  // 完成
	Cancel   SysTaskStatus = 3  // 取消
)

type TaskHTTPMethod string

const (
	TaskHTTPMethodGet  TaskHTTPMethod = "GET"
	TaskHTTPMethodPost TaskHTTPMethod = "POST"
)
