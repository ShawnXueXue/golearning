package main

import (
	"fmt"
	"net/http"
)

var pf = fmt.Printf
var pl = fmt.Println

func main() {
	host := "github.com"

	// 1
	url1 := "http://" + host
	pf("Send request to %q with method GET ...\n", url1)
	resp1, err := http.Get(url1)
	if err != nil {
		pf("request sending error: %v\n", err)
		return
	}
	defer resp1.Body.Close()
	line1 := resp1.Proto + " " + resp1.Status
	pf("The first line of response:\n%s\n", line1)
	pl()

	// 2
	url2 := "http://" + host + "/ShawnXueXue"
	pf("Send request to %q with method GET ...\n", url2)
	client := http.Client{}
	resp2, err := client.Get(url2)
	if err != nil {
		pf("request sending error: %v\n", err)
		return
	}
	defer resp2.Body.Close()
	line2 := resp2.Proto + " " + resp2.Status
	pf("The first line of response:\n%s\n", line2)
}
