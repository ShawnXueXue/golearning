package main

import "fmt"

func init() {
	fmt.Println("init func.")
}

const (
	branchPageFlag   = 0x01
	leafPageFlag     = 0x02
	metaPageFlag     = 0x04
	freelistPageFlag = 0x10
)

func main() {
	//fmt.Println("main func.")
	//fmt.Println(add(1, 2))
	//fmt.Println(add(1, 2, 3))
	//fmt.Println(add([]int{1, 2, 3}...))
	fmt.Println(typ(33))
	fmt.Println(0x7FFFFFF)
}

func typ(flags uint16) string {
	if (flags & branchPageFlag) != 0 {
		return "branch"
	} else if (flags & leafPageFlag) != 0 {
		return "leaf"
	} else if (flags & metaPageFlag) != 0 {
		return "meta"
	} else if (flags & freelistPageFlag) != 0 {
		return "freelist"
	}
	return fmt.Sprintf("unknown<%02x>", flags)
}

func add(args ...int) (sum int) {
	for _, arg := range args {
		sum += arg
	}
	return
}
