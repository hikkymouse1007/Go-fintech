package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("Init!")
}

func buzz()  {
	fmt.Println("Buzz")
}

func main() {
	//buzz()
	fmt.Println("hello world", time.Now())
	fmt.Println(user.Current())
}
