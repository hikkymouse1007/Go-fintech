package main

import "fmt"

func main() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 1, t)
	fmt.Printf("%T %v %t\n", f, 0, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || false)
}
