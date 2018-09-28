package main
import ("go/format"
	"io/ioutil")
var filePath="format.go"
var formatFilePath="formatted.go"
func main(){
	bytes, e:=ioutil.ReadFile(filePath)
	if e != nil{
		panic(e)
	}
source, i := format.Source(bytes)
	if i != nil{
		panic(i)
	}
	we:=ioutil.WriteFile(formatFilePath, source,0644)
	if we!=nil{
		panic(we)
	}}