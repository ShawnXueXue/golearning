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
	profileName      = "blockprofile.out"
	blockProfileRate = 2
	debug            = 0
)

func main() {
	file, e := common.CreateFile("", profileName)
	if e != nil {
		fmt.Printf("block profile creation error: %v\n", e)
		return
	}
	defer file.Close()
	startBlockProfile()
	if e = common.Execute(op.BlockProfile, 10); e != nil {
		fmt.Printf("execute error: %v\n", e)
		return
	}
	if e = stopBlockProfile(file); e != nil {
		fmt.Printf("block profile stop error: %v\n", e)
		return
	}
}

func startBlockProfile() {
	runtime.SetBlockProfileRate(blockProfileRate)
}

func stopBlockProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.Lookup("block").WriteTo(f, debug)
}
