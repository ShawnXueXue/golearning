package main

import (
	"errors"
	"fmt"
	"golearning/puzzlers/pprof/common"
	"golearning/puzzlers/pprof/common/op"
	"os"
	"runtime/pprof"
)

var (
	profileName = "cpuprofile.out"
)

/**
go get -u github.com/google/pprof
pprof -http=:8080 cpuprofile.out

如果是service类型的程序, 还可以import _ "net/http/pprof", 访问http://xxxx/debug/pprof/
go tool pprof -inuse_space http://xxxx/debug/pprof/heap
	1. top10来查看正在使用的对象较多的10个函数入口
go tool pprof -alloc_space http://xxxx/debug/pprof/heap
	1. top来查看累积分配内存较多的一些函数调用
	2. top -cum来查找，-cum的意思就是，将函数调用关系 中的数据进行累积，比如A函数调用的B函数，则B函数中的内存分配量也会累积到A上面，这样就可以很容易的找出调用链。
	3. web查看图片

如果程序已经没有反应了, 可以 kill -SIGQUIT <pid>, 可执行程序统计目录下生成core.pid文件, 使用delve main core.pid 分析dump
需要预先做两件事: ulimit -c unlimited; GOTRACEBACK=crash ./main
*/
func main() {
	file, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("CPU profile creation error: %v\n", err)
		return
	}
	defer file.Close()
	if err = startCPUProfile(file); err != nil {
		fmt.Printf("CPU profile start error: %v\n", err)
		return
	}
	if err = common.Execute(op.CPUProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
	stopCPUProfile()
}

func startCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}
