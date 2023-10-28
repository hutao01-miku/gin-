package main

import (
	"fmt"
	"os"
	"time"
)

// 了解就行
func main() {
	//获得当前正在运行的进程的pid
	fmt.Println(os.Getpid())
	//获得当前进程的父进程的pid
	fmt.Println(os.Getppid())
	//设置新进程的属性
	atrr := &os.ProcAttr{
		//files属性是进程的文件描述符，这里是继承当前进程的文件描述符
		//前三个分别是标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		//新进程的环境变量
		Env: os.Environ(),
	}
	//启动新进程
	p, err := os.StartProcess("F:\\app\\typora\\Typora\\Typora.exe", []string{"F:\\app\\typora\\Typora\\Typora.exe", "F:\\huancun\\go\\awesomeProject1\\a.txt"}, atrr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Pid)
	//通过进程的pid获得进程对象
	p, err = os.FindProcess(p.Pid)
	fmt.Println(p.Pid)
	//等待10秒，执行函数
	time.AfterFunc(time.Second*10, func() {
		//杀死进程(向p进程发出退出消息)
		p.Signal(os.Kill)
	})
	//等待p 进程退出,返回进程状态
	ps, _ := p.Wait()
	fmt.Println(ps.String())
}
