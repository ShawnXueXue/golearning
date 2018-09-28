package main

import "fmt"

var container = []string{"zero", "one", "two"}

func getElement(container interface{}) (elem string, err error) {
	switch t := container.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", container)
		return
	}
	return
}

func main() {
	container := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Printf("The element is %q\n", container[1])
	fmt.Printf("The type is %T\n", container)

	// 1
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)
	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	fmt.Printf("The element is %q. container type: %T\n", container[1], container)

	// 2
	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q. container type: %T\n", elem, container)
}
