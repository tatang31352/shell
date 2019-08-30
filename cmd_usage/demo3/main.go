package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err error
	output []byte
}

func main() {
	//执行1个cmd,让它在一个协成里去执行,让它执行2秒:sellp 2;echo hello;
	//1秒的时候,我们杀死cmd
	var(
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)

	//创建一个结果队列
	resultChan = make(chan *result,1000)

	ctx,cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var(
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx,"D:\\cygwin64\\bin\\bash.exe","-c","echo hello;")

		//执行任务,捕获输出
		output,err = cmd.CombinedOutput()

		//把任务输出结果,传给main协成
		resultChan <- &result{
			err:err,
			output:output,
		}

	}()

	//继续往下走
	time.Sleep(1 * time.Second)

	//取消上下文
	cancelFunc()

	res = <- resultChan

	fmt.Println(res.err,string(res.output))


}
