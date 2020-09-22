package main

import "fmt"

//可変長引数

func main() {
	f := 1.11
	int_f := int(f)
	fmt.Printf("%T, %v\n", int_f, int_f)

	m := map[string]int{
		"Mike":20,
		"Nancy":24,
		"Messi":30,
	}

	fmt.Printf("%T %v", m, m )
}
