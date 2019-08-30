package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	//生成Cmd
	cmd = exec.Command("D:\\cygwin64\\bin\\bash.exe","-c","sleep 1;ls -l;sleep 5;echo 222;")

	//执行了命令,捕获了子进程的输（pipe）
	if output,err = cmd.CombinedOutput();err != nil{
		fmt.Println(err)
		return
	}

	//打印子进程的输出
	fmt.Println(string(output))

}
