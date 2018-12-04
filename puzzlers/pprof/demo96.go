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
