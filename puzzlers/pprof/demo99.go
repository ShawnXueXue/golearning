package main

import "runtime"

func createDir() (string, error) {

}

func prepare() {
	runtime.MemProfileRate = 8
	runtime.SetBlockProfileRate(2)
}
