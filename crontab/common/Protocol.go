package common

import (
	"context"
	"encoding/json"
	"github.com/gorhill/cronexpr"
	"time"
)

// 定时任务
type Job struct {
	Name string `json:"name"`	//  任务名
	Command string	`json:"command"` // shell命令
	CronExpr string	`json:"cronExpr"`	// cron表达式
}

// 任务调度计划
type JobSchedulePlan struct {
	Job *Job	// 要调度的任务信息
	Expr *cronexpr.Expression	// 解析好的cronexpr表达式
	NextTime time.Time	// 下次调度时间
}

// 任务执行状态
type JobExecuteInfo struct {
	Job *Job // 任务信息
	PlanTime time.Time // 理论上的调度时间
	RealTime time.Time // 实际的调度时间
	CancelCtx context.Context // 任务command的context
	CancelFunc context.CancelFunc//  用于取消command执行的cancel函数
}

// 任务执行日志
type JobLog struct {
	JobName string `json:"jobName" bson:"jobName"` // 任务名字
	Command string `json:"command" bson:"command"` // 脚本命令
	Err string `json:"err" bson:"err"` // 错误原因
	Output string `json:"output" bson:"output"`	// 脚本输出
	PlanTime int64 `json:"planTime" bson:"planTime"` // 计划开始时间
	ScheduleTime int64 `json:"scheduleTime" bson:"scheduleTime"` // 实际调度时间
	StartTime int64 `json:"startTime" bson:"startTime"` // 任务执行开始时间
	EndTime int64 `json:"endTime" bson:"endTime"` // 任务执行结束时间
}

// 日志批次
type LogBatch struct {
	Logs []interface{}	// 多条日志
}

// 任务日志过滤条件
type JobLogFilter struct {
	JobName string `bson:"jobName"`
}

// 任务日志排序规则
type SortLogByStartTime struct {
	SortOrder int `bson:"startTime"`	// {startTime: -1}
}

// HTTP接口应答
type Response struct {
	Errno int `json:"errno"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}


// 应答方法
func BuildResponse(errno int, msg string, data interface{}) (resp []byte, err error) {
	// 1, 定义一个response
	var (
		response Response
	)

	response.Errno = errno
	response.Msg = msg
	response.Data = data

	// 2, 序列化json
	resp, err = json.Marshal(response)
	return
}