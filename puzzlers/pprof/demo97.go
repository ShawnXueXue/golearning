package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golearning/puzzlers/pprof/common"
	"golearning/puzzlers/pprof/common/op"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	profileName    = "memprofile.out"
	memProfileRate = 8
)

func main() {
	file, e := common.CreateFile("", profileName)
	if e != nil {
		fmt.Printf("memory creation error: %v\n", e)
		return
	}
	defer file.Close()
	startMemProfile()
	if e = common.Execute(op.MemProfile, 10); e != nil {
		fmt.Printf("execute error: %v\n", e)
		return
	}
	if e = stopMemProfile(file); e != nil {
		fmt.Printf("memory profile stop error: %v\n", e)
		return
	}
}

func startMemProfile() {
	runtime.MemProfileRate = memProfileRate
}

func stopMemProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.WriteHeapProfile(f)
}
