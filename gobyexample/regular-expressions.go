package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	pl(match) //true

	r, _ := regexp.Compile("p([a-z]+)ch")

	pl(r.MatchString("peach"))                                // true
	pl(r.FindString("peach punch"))                           // peach
	pl(r.FindStringIndex("peach punch"))                      // [0 5]
	pl(r.FindStringSubmatch("peach punch"))                   // [peach ea]
	pl(r.FindStringSubmatchIndex("peach punch"))              //[0 5 1 3]
	pl(r.FindAllString("peach punch pinch", -1))              //[peach punch pinch]
	pl(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) //[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
	pl(r.FindAllString("peach punch pinch", 2))               //[peach punch]  2为限制显示次数
	pl(r.Match([]byte("peach")))                              //true

	r = regexp.MustCompile("p([a-z]+)ch")
	pl(r) //p([a-z]+)ch

	pl(r.ReplaceAllString("a peach", "<fruit>")) //a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	pl(string(out)) //a PEACH
}
