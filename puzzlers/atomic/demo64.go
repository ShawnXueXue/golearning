package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync/atomic"
)

func main() {
	// 1
	var box atomic.Value
	fmt.Println("copy box to box2.")
	box2 := box
	v1 := []int{1, 2, 3}
	fmt.Printf("Store %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())
	fmt.Println()

	// 2
	v2 := "123"
	fmt.Printf("Store %q to box2.\n", v2)
	box2.Store(v2)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %q.\n", box2.Load())
	fmt.Println()

	// 3
	fmt.Println("copy box to box3.")
	box3 := box
	fmt.Printf("The value load from box3 is %v.\n", box3.Load())
	v3 := 123
	fmt.Printf("Store %d to box3.\n", v3)
	//box3.Store(v3) // panic: store of inconsistently typed value into Value
	_ = box3
	fmt.Println()

	// 4
	var box4 atomic.Value
	v4 := errors.New("something wrong")
	fmt.Printf("Store an error message %q to box4.\n", v4)
	box4.Store(v4)
	v41 := io.EOF
	fmt.Println("Store a value of the same type to box4.")
	box4.Store(v41)
	v42, ok := interface{}(&os.PathError{}).(error)
	if ok {
		fmt.Printf("Store a value of type %T that implements error interface to box4.\n", v42)
		//box4.Store(v42) // panic: store of inconsistently typed value into Value
	}
	fmt.Println()

	// 5
	box5, err := NewAtomicValue(v4)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("The legal type in box5 is %s.\n", box5.TypeOfValue())
	fmt.Println("Store a value of the same type to box5.")
	err = box5.Store(v41)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("Store a value of type %T that implements error interface to box5.\n", v42)
	// ???
	box5.Store(v42)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Println()

	// 6
	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	fmt.Printf("Store %v to box6.\n", v6)
	box6.Store(v6)
	v6[1] = 4
	fmt.Printf("The value load from box6 is %v\n", box6.Load())
	// right way
	v6 = []int{1, 2, 3}
	store := func(v []int) {
		replace := make([]int, len(v))
		copy(replace, v)
		box6.Store(replace)
	}
	fmt.Printf("Store %v to box6.\n", v6)
	store(v6)
	v6[1] = 4
	fmt.Printf("The value load from box6 is %v\n", box6.Load())
}

type atomicValue struct {
	v atomic.Value
	t reflect.Type
}

func NewAtomicValue(example interface{}) (*atomicValue, error) {
	if example == nil {
		return nil, errors.New("atomic value: nil example")
	}
	return &atomicValue{
		t: reflect.TypeOf(example),
	}, nil
}

func (av *atomicValue) Store(v interface{}) error {
	if v == nil {
		return errors.New("atomic value: nil value")
	}
	t := reflect.TypeOf(v)
	if t != av.t {
		return fmt.Errorf("atomic value: wrong type: %s", t)
	}
	av.v.Store(v)
	return nil
}

func (av *atomicValue) Load() interface{} {
	return av.v.Load()
}

func (av *atomicValue) TypeOfValue() reflect.Type {
	return av.t
}
