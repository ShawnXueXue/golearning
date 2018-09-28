package main

import (
	"fmt"
)

var pln = fmt.Println
var pf = fmt.Printf

func main() {
	array1 := [3]string{"a", "b", "c"}
	pf("The array: %v\n", array1)
	pf("The input addr before func:      %p\n", &array1)
	array2 := modifyArray(array1)
	pf("The result addr out modifyArray: %p\n", &array2)
	pf("The modified array: %v\n", array2)
	pf("The original array: %v\n", array1)
	pln()

	array3 := &[3]string{"a", "b", "c"}
	pf("The array: %v\n", array3)
	pf("The input addr before func:        %p\n", &array3)
	pf("The input value before func:       %p\n", array3)
	array4 := modifyArrayP(array3)
	pf("The result addr out modifyArrayP:  %p\n", &array4)
	pf("The result valur out modifyArrayP: %p\n", array4)
	pf("The modified array: %v\n", array4)
	pf("The original array: %v\n", array4)
	pln()

	slice1 := []string{"x", "y", "z"}
	pf("The slice: %v\n", slice1)
	pf("The input addr before func:      %p\n", &slice1)
	slice2 := modifySlice(slice1)
	pf("The result addr out modifySlice: %p\n", &slice2)
	pf("The modified slice: %v\n", slice2)
	pf("The original slice: %v\n", slice1)
	pln()

	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	pf("The complec array: %v\n", complexArray1)
	pf("The input addr before func:             %p\n", &complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	pf("The result addr out modifyComplexArray: %p\n", &complexArray2)
	pf("The modified complex array: %v\n", complexArray2)
	pf("The original complex array: %v\n", complexArray1)
}

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	pf("The result addr in modifyArray:  %p\n", &a)
	return a
}

func modifyArrayP(a *[3]string) *[3]string {
	a[1] = "x"
	pf("The result addr in modifyArrayP:   %p\n", &a)
	pf("The result valur in modifyArrayP:  %p\n", a)
	return a
}

func modifySlice(a []string) []string {
	a[1] = "i"
	pf("The result addr in modifySlice:  %p\n", &a)
	return a
}

func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"
	a[2] = []string{"o", "p", "q"}
	pf("The result addr in modifyComplexArray:  %p\n", &a)
	return a
}
