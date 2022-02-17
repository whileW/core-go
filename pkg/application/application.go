package application

import (
	"fmt"
	"github.com/whileW/core-go/pkg/conf"
	"github.com/whileW/core-go/pkg/flag"
	"github.com/whileW/core-go/pkg/log"
	"github.com/whileW/core-go/pkg/util/xgo"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Service interface {
	StartUp() error
	Stop() error
}
var (
	StartTime = time.Now()
	initializeFns = []func() error {
		flag.Initialize,
		conf.Initialize,
		log.Initialize,
	}
	initOnce = sync.Once{}
	programs = []Service{}
)
func initialize() {
	initOnce.Do(func() {
		fmt.Println("开始初始化相关启动项 ...")
		st := time.Now()
		if err := xgo.SerialUntilError(initializeFns...)();err != nil{
			panic(err)
		}
		fmt.Println("初始化成功，耗时： "+fmt.Sprint(time.Now().Sub(st)))
	})
	return
}
func StartUp(p Service)  {
	initialize()
	programs = append(programs, p)
	go func() {
		if err := p.StartUp();err != nil {
			panic(err)
		}
	}()
}
func KeepAlive()  {
	/*
					监听系统信号
		SIGHUP = 终端控制进程结束(终端连接断开)
		SIGQUIT = 用户发送QUIT字符(Ctrl+/)触发
		SIGTERM = 结束程序(可以被捕获、阻塞或忽略)
		SIGINT = 用户发送INTR字符(Ctrl+C)触发
	*/

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-c:
			for _,p := range programs {
				if err := p.Stop();err != nil {
					fmt.Println(err.Error())
				}
			}
			os.Exit(0)
		}
	}
}

func ResetInitializeFns(fns ...func()error) {
	initializeFns = fns
}
func AppendInitializeFns(fns ...func()error) {
	initializeFns = append(initializeFns, fns...)
}