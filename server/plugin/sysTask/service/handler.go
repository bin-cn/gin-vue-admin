package service

import (
	"bytes"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	taskGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/model"
	customclient "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/utils/httpclient"
	"go.uber.org/zap"
	"net/http"
	"os/exec"
)

// Handler 任务处理程序
type Handler interface {
	Run(taskModel model.SysTask, sysTaskLogId int) (string, error)
}

// ShellHandler Shell任务
type ShellHandler struct{}

// Run 执行Shell任务
func (h *ShellHandler) Run(taskModel model.SysTask, sysTaskLogId int) (result string, err error) {
	global.GVA_LOG.Info("==========执行内部Shell的生成逻辑==========")
	shellPath := taskModel.InvokeTarget

	// 创建命令
	cmd := exec.Command("sh", "-c", shellPath)

	// 获取标准输出和标准错误管道
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 启动命令
	if err := cmd.Start(); err != nil {
		logMessage := fmt.Sprintf("启动命令失败: %s", err)
		global.GVA_LOG.Error(logMessage)
		return "", err
	}

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		logMessage := fmt.Sprintf("命令执行失败: %s\n标准错误: %s", err, stderr.String())
		global.GVA_LOG.Error(logMessage)
		return "", err
	}

	// 获取命令的输出结果
	output := stdout.String()
	global.GVA_LOG.Info("命令输出: " + output)
	return output, nil
}

// HTTPHandler HTTP任务
type HTTPHandler struct{}

func (h *HTTPHandler) Run(taskModel model.SysTask, sysTaskLogId int) (result string, err error) {
	if taskModel.Timeout <= 0 || taskModel.Timeout > taskGlobal.HttpExecTimeout {
		taskModel.Timeout = taskGlobal.HttpExecTimeout
	}
	var resp customclient.ResponseWrapper
	// 返回状态码非200，均为失败
	switch taskModel.HttpMethod {
	case taskGlobal.TaskHTTPMethodGet:
		resp = customclient.Get(taskModel.InvokeTarget, taskModel.Timeout)
	case taskGlobal.TaskHTTPMethodPost:
		resp = customclient.PostJson(taskModel.InvokeTarget, taskModel.Args, taskModel.Timeout)
	default:
		global.GVA_LOG.Info("暂时还没有其他method类型，可自行添加")
	}
	if resp.StatusCode != http.StatusOK {
		global.GVA_LOG.Error("[ERROR] http request failed! ", zap.Error(err))
		return resp.Body, fmt.Errorf("HTTP状态码非200-->%d", resp.StatusCode)
	}

	return resp.Body, err
}

// InternalFuncHandler 内部函数任务
type InternalFuncHandler struct{}

func (h *InternalFuncHandler) Run(taskModel model.SysTask, sysTaskLogId int) (result string, err error) {
	global.GVA_LOG.Info("==========执行内部函数的生成逻辑==========")

	var obj = JobList[taskModel.InvokeTarget]
	if obj == nil {
		global.GVA_LOG.Warn("[Job] ExecJob Run job nil, invokeTarget: " + taskModel.InvokeTarget)
		return
	}
	err = CallExec(obj.(JobExec), taskModel.Args)
	if err != nil {
		// 如果失败暂停一段时间重试
		global.GVA_LOG.Error(" [ERROR] mission failed! ", zap.Error(err))
		result = err.Error()
	} else {
		result = "success"
	}
	return result, err
}
