package main

import "fmt"

func goroutine1(s []int, c chan int)  {
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum  //channelにsumを送る, returnの代わり
}


func goroutine2(s []int, c chan int)  {
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum  //channelにsumを送る
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // アンバッファ 15, 15
	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
