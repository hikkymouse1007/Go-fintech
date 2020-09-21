package main

import "fmt"

func main() {
	// ASCIIの72, 73
	b := []byte{72, 73}
	fmt.Println(b)
	// 72:H, 73:I
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}
