## GVA 定时任务管理配置好
#### 开发者：dk.jin

### 使用步骤

#### 1. 前往GVA主程序下的initialize/router.go 在Routers 方法最末尾按照你需要的及安全模式添加本插件
	PluginInit(PrivateGroup, systask.CreateSysTaskPlug())

### 2. 前往GVA主程序下的initialize/time.go在Timer()方法下面添加定时任务的启动程序,保证系统启动时将定时任务启动起来
    func Timer() {
        service.ServiceGroupApp.SysTaskService.Initialize()
    }

#### 2-1 全局配置结构体说明
   
#### 2-2 入参结构说明
    TaskName       string               `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:32;"`                              //任务名称
	TaskGroup      string               `json:"taskGroup" form:"taskGroup" gorm:"column:task_group;comment:任务分组;size:255;"`                          //任务分组
	CronExpression string               `json:"cronExpression" form:"cronExpression" gorm:"column:cron_expression;comment:cron_expression;size:64;"` //cron_expression
	InvokeTarget   string               `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:;size:255;"`                     //invokeTarget
	Args           string               `json:"args" form:"args" gorm:"column:args;comment:;size:255;"`                                              //args
	Timeout        int                  `json:"timeout" form:"timeout" gorm:"column:timeout;comment:任务执行超时时间(单位秒),0不限制;"`                            //任务执行超时时间(单位秒),0不限制
	Multi          int                  `json:"multi" form:"multi" gorm:"column:multi;comment:是否允许多实例运行;"`                                           //是否允许多实例运行
	RetryTimes     int                  `json:"retryTimes" form:"retryTimes" gorm:"column:retry_times;comment:重试次数;"`                                //重试次数
	RetryInterval  int                  `json:"retryInterval" form:"retryInterval" gorm:"column:retry_interval;comment:重试间隔时间;"`                     //重试间隔时间
	Tag            string               `json:"tag" form:"tag" gorm:"column:tag;comment:;size:32;"`                                                  //tag
	Remark         string               `json:"remark" form:"remark" gorm:"column:remark;comment:;size:100;"`                                        //remark
	Status         global.SysTaskStatus `json:"status" form:"status" gorm:"column:status;comment:状态 1:正常 0:停止;"`
#### 2-3 SQL文件
    先添加SQL文件，/plugin/sys_task/sql/sys_task.sql
#### 2-4 配置sys_apis的接口
    INSERT INTO `sys_apis`(path,description,api_group,method )VALUES 
    ('/sysTask/sysTask/createSysTask','新增sysTask表','sysTask表','POST'),
    ('/sysTask/sysTask/deleteSysTask','删除sysTask表','sysTask表','DELETE'),
    ('/sysTask/sysTask/deleteSysTaskByIds','批量删除sysTask表','sysTask表','DELETE'),
    ('/sysTask/sysTask/updateSysTask','更新sysTask表','sysTask表','PUT'),
    ('/sysTask/sysTask/findSysTask','根据ID获取sysTask表','sysTask表','GET'),
    ('/sysTask/sysTask/getSysTaskList','获取sysTask表列表','sysTask表','GET'),
    ('/sysTask/sysTaskLog/createSysTaskLog','新增sysTaskLog表','sysTaskLog表','POST'),
    ('/sysTask/sysTaskLog/deleteSysTaskLog','删除sysTaskLog表','sysTaskLog表','DELETE'),
    ('/sysTask/sysTaskLog/deleteSysTaskLogByIds','批量删除sysTaskLog表','sysTaskLog表','DELETE'),
    ('/sysTask/sysTaskLog/updateSysTaskLog','更新sysTaskLog表','sysTaskLog表','PUT'),
    ('/sysTask/sysTaskLog/findSysTaskLog','根据ID获取sysTaskLog表','sysTaskLog表','GET'),
    ('/sysTask/sysTaskLog/getSysTaskLogList','获取sysTaskLog表列表','sysTaskLog表','GET'),
    ('/sysTask/sysTask/runOneSysTask','任务运行一次','sysTask表','GET'),
    ('/sysTask/sysTask/enableSysTask','激活任务','sysTask表','GET'),
    ('/sysTask/sysTask/disableSysTask','暂停任务','sysTask表','GET'),
    ('/sysTask/sysTask/removeSysTask','删除任务','sysTask表','GET');

### 3. 方法API

    sysTask/sysTask/findSysTask，根据ID获取sysTask表
    sysTask/sysTask/getSysTaskList, 获取sysTask表列表
    sysTask/sysTask/runOneSysTask, 运行一次任务
    sysTask/sysTask/enableSysTask， 激活任务
    sysTask/sysTask/disableSysTask， 暂停任务
    sysTask/sysTask/removeSysTask， 删除任务
    
    sysTask/sysTaskLog/findSysTaskLog， 根据ID获取sysTaskLog表
    sysTask/sysTaskLog/getSysTaskLogList， 获取sysTaskLog表列表

### 4. 可直接调用的接口

### 5. 安装cron sdk
    go get github.com/jakecoffman/cron


    
   
