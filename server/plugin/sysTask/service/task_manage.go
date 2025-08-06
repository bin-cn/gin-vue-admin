package service

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	globalTask "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/utils"
	"github.com/jakecoffman/cron"
	"go.uber.org/zap"
	"strconv"
	"sync"
	"time"
)

var (
	// 定时任务调度管理器
	serviceCron *cron.Cron

	// 同一任务是否有实例处于运行中
	runInstance Instance

	// 任务计数-正在运行的任务
	taskCount TaskCount

	// 并发队列, 限制同时运行的任务数量
	concurrencyQueue ConcurrencyQueue
)

// ConcurrencyQueue 并发队列
type ConcurrencyQueue struct {
	queue chan struct{}
}

func (cq *ConcurrencyQueue) Add() {
	cq.queue <- struct{}{}
}

func (cq *ConcurrencyQueue) Done() {
	<-cq.queue
}

// TaskCount 任务计数
type TaskCount struct {
	wg   sync.WaitGroup
	exit chan struct{}
}

func (tc *TaskCount) Add() {
	tc.wg.Add(1)
}

func (tc *TaskCount) Done() {
	tc.wg.Done()
}

func (tc *TaskCount) Exit() {
	tc.wg.Done()
	<-tc.exit
}

func (tc *TaskCount) Wait() {
	tc.Add()
	tc.wg.Wait()
	close(tc.exit)
}

// Instance 任务ID作为Key
type Instance struct {
	m sync.Map
}

// 是否有任务处于运行中
func (i *Instance) has(key int) bool {
	_, ok := i.m.Load(key)

	return ok
}

func (i *Instance) add(key int) {
	i.m.Store(key, struct{}{})
}

func (i *Instance) done(key int) {
	i.m.Delete(key)
}

// Initialize 初始化任务, 从数据库取出所有任务, 添加到定时任务并运行
func (sysTaskService *SysTaskService) Initialize() {
	InitJob()

	serviceCron = cron.New()
	serviceCron.Start()
	concurrencyQueue = ConcurrencyQueue{queue: make(chan struct{}, 500)}
	taskCount = TaskCount{sync.WaitGroup{}, make(chan struct{})}
	go taskCount.Wait()
	global.GVA_LOG.Info("开始初始化定时任务")
	taskNum := 0
	taskList, err := SysTaskServiceApp.ActiveSysTaskList()
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败", zap.Error(err))
	}
	for _, item := range taskList {
		sysTaskService.Add(item)
		taskNum++
	}
	global.GVA_LOG.Info("初始化任务完成", zap.Int("taskNum", taskNum))
}

// RemoveAndAdd 删除任务后添加
func (sysTaskService *SysTaskService) RemoveAndAdd(taskModel model.SysTask) {
	//task.Remove(taskModel.Id)
	sysTaskService.Add(taskModel)
}

// Remove 删除任务
func (sysTaskService *SysTaskService) Remove(id string) {
	serviceCron.RemoveJob(id)
}

// Add 添加任务
func (sysTaskService *SysTaskService) Add(taskModel model.SysTask) {
	taskFunc := CreateJob(taskModel)
	if taskFunc == nil {
		fmt.Println("创建任务处理Job失败,不支持的任务协议#")
		return
	}

	cronName := strconv.Itoa(taskModel.ID)
	err := utils.PanicToError(func() {
		serviceCron.AddFunc(taskModel.CronExpression, taskFunc, cronName)
	})
	if err != nil {
		global.GVA_LOG.Error("添加任务到调度器失败", zap.Any("err", err))
	}
}

// ChangeStatus 改变任务状态
func (sysTaskService *SysTaskService) ChangeStatus(idStr string, status globalTask.SysTaskStatus) string {

	id, _ := strconv.Atoi(idStr)
	err := sysTaskService.UpdateTaskStatus(id, status)
	if err != nil {
		return "更新任务状态失败"
	}

	if status == globalTask.Enabled {
		sysTaskService.AddTaskToTimer(idStr)
	} else {
		sysTaskService.Remove(idStr)
	}
	return "任务状态修改成功"
}

// Run 直接运行任务
func (sysTaskService *SysTaskService) Run(taskModel model.SysTask) {
	go CreateJob(taskModel)()
}

// CreateJob  创建任务处理Job
func CreateJob(taskModel model.SysTask) cron.FuncJob {
	handler := CreateHandler(taskModel)
	if handler == nil {
		return nil
	}
	taskFunc := func() {
		taskCount.Add()
		defer taskCount.Done()

		taskLogId := beforeExecJob(taskModel)
		if taskLogId <= 0 {
			return
		}

		if taskModel.Multi == 0 {
			runInstance.add(taskModel.ID)
			defer runInstance.done(taskModel.ID)
		}

		concurrencyQueue.Add()
		defer concurrencyQueue.Done()
		global.GVA_LOG.Info(fmt.Sprintf("开始执行任务命令-%s", taskModel.InvokeTarget))
		taskResult := execJob(handler, taskModel, taskLogId)
		global.GVA_LOG.Info(fmt.Sprintf("任务执行完成命令-%s", taskModel.InvokeTarget))
		afterExecJob(taskResult, taskLogId)
	}

	return taskFunc
}

// CreateHandler 创建任务处理Handler
func CreateHandler(taskModel model.SysTask) Handler {
	var handler Handler = nil
	switch taskModel.Tag {
	case "httpTask":
		handler = new(HTTPHandler)
	case "internalFuncTask":
		handler = new(InternalFuncHandler)
	case "shellTask":
		handler = new(ShellHandler)
	}
	return handler
}
func (sysTaskService *SysTaskService) AddTaskToTimer(id string) {
	// 获取任务详情
	task, err := sysTaskService.GetSysTask(id)
	if err != nil {
		global.GVA_LOG.Error("获取任务详情失败", zap.Error(err))
		return
	}
	sysTaskService.RemoveAndAdd(task)
}

// 任务前置操作
func beforeExecJob(taskModel model.SysTask) (taskLogId int) {
	if taskModel.Multi == 0 && runInstance.has(taskModel.ID) {
		taskLogId, _ := SysTaskLogServiceApp.CreateTaskLogByTask(taskModel, globalTask.Cancel)
		return taskLogId
	}
	taskLogId, err := SysTaskLogServiceApp.CreateTaskLogByTask(taskModel, globalTask.Running)
	if err != nil {
		global.GVA_LOG.Error("任务开始执行#写入任务日志失败", zap.Error(err))
		return
	}
	return taskLogId
}

// 任务执行后置操作
func afterExecJob(taskResult model.SysTaskResult, taskLogId int) {
	_, err := SysTaskLogServiceApp.UpdateTaskLogByTaskResult(taskLogId, taskResult)
	if err != nil {
		global.GVA_LOG.Error("任务结束#更新任务日志失败-", zap.Error(err))
	}
}

// 执行具体任务
func execJob(handler Handler, taskModel model.SysTask, sysTaskLogId int) model.SysTaskResult {
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Error("panic#task#service.go.bak:execJob#")
		}
	}()
	// 默认只运行任务一次
	var execTimes int8 = 1
	if taskModel.RetryTimes > 0 {
		execTimes += int8(taskModel.RetryTimes)
	}

	var taskLog model.SysTaskLog
	global.GVA_DB.Where("id =?", sysTaskLogId).First(&taskLog)
	var i int8 = 0
	var output string
	var err error
	for i < execTimes {
		output, err = handler.Run(taskModel, sysTaskLogId)
		if err == nil {
			return model.SysTaskResult{Result: output, Err: err, SysTaskLog: taskLog}
		}
		i++
		if i < execTimes {
			global.GVA_LOG.Error(fmt.Sprintf("任务执行失败#任务id-%d#重试第%d次#输出-%s#错误-%s", taskModel.ID, i, output), zap.Error(err))
			if taskModel.RetryInterval > 0 {
				time.Sleep(time.Duration(taskModel.RetryInterval) * time.Second)
			} else {
				// 默认重试间隔时间，每次递增1分钟
				time.Sleep(time.Duration(i) * time.Minute)
			}
		}
	}

	return model.SysTaskResult{Result: output, Err: err, SysTaskLog: taskLog}
}
