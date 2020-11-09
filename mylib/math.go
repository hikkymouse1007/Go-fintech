/*
mylib is my special library
*/
package mylib

import "fmt"

type Person2 struct {
	//Name
	Name string
	//Age
	Age int
}

func (p *Person2) Say(){
	fmt.Println("Person2")
}

// Average returns the average of a series of numbers
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
