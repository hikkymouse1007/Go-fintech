package main

import "fmt"

func main()  {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch) //rangeで呼び出すときはcloseが必要

	for c := range ch{
		fmt.Println(c)
	}
}
