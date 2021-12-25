package main

import (
	"fmt"
	"jsonmaker/function"
)

func main() {

	var st string
	fmt.Println("Write a filename")
	fmt.Scan(&st)
	var fvalue []string = function.GetFileText(st)
	var result = function.TextToJson(fvalue)
	function.CreateJsonFile(result)
	//function.CreateJson(st)
}
