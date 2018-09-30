package consttest

import (
	"fmt"
	"reflect"
	"unsafe"
)

const a = "1"
const leafPageElementSize = int(unsafe.Sizeof(leafPageElement{}))

func t1() {
	fmt.Println(pageHeaderSize)
	fmt.Println(leafPageElementSize)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(b)
}
