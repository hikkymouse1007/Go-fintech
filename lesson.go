package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Hello" + "World!")
	fmt.Println("Hello World"[0])         //ASCII表示
	fmt.Println(string("Hello World"[0])) //typecasting
	var s string = "Hello World"
	//fmt.Println(strings.Replace(s, "H", "X", 1))
	s = strings.Replace(s, "H", "X", 1) //sの上書き
	fmt.Println(s)
	fmt.Println(`test
				test
test`)
	fmt.Println("\"")
	fmt.Println(`"`)
}
