package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

//代表一个任务
type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time //expr.Next(now)
}

func main() {
	//需要有1个调度协成,它定时检查所有的Cron任务,谁过期了就执行谁

	var (
		cronJon       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob //key:任务的名字
	)

	scheduleTable = make(map[string]*CronJob)

	//1,我们定义2个cronJob
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJon = &CronJob{
		expr,
		expr.Next(now),
	}
	//任务注册到调度表
	scheduleTable["job1"] = cronJon

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJon = &CronJob{
		expr,
		expr.Next(now),
	}
	//任务注册到调度表
	scheduleTable["job2"] = cronJon
	//启动一个调度协程
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		//定时检查一下任务调度表
		for {
			now = time.Now()
			for jobName, cronJon = range scheduleTable {
				//判断是否过期
				if cronJon.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					//启动一个协程,执行这个任务
					go func(jobName string) {
						fmt.Println("执行", jobName)
					}(jobName)

					//计算下一次调度时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "下次执行时间:", cronJob.nextTime)
				}
			}

			//睡眠100毫秒
			//time.Sleep(100 * time.Millisecond)
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:
			}
		}
	}()

	time.Sleep(100 * time.Second)
}
