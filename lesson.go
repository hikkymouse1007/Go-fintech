package main

import (
	"fmt"
	"sync"
	"time"
)


func producer(ch chan int, i int)  {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup)  {
	for i := range ch{

		fmt.Println("process", i * 1000)
		wg.Done() //closeしないと次のchannelを待っていて、for文が実行待ちになってしまう
	}
	fmt.Println("#############") //closeしないとここは実行されない
}

func main()  {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1) //channelを生成
		go producer(ch, i)
	}

	go consumer(ch, &wg)
	wg.Wait()
	close(ch) // 10回実行後もconsumerのgoroutineがchannelを取ろうと待ってしまうためにcloseする
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
