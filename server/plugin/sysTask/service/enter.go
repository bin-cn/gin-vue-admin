package service

type ServiceGroup struct {
	SysTaskService
	SysTaskLogService
}

var ServiceGroupApp = new(ServiceGroup)
