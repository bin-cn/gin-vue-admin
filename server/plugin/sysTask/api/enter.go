package api

type ApiGroup struct {
	SysTaskApi
	SysTaskLogApi
}

var SysTaskGroupApp = new(ApiGroup)
