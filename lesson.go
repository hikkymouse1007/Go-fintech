package main

import (
	"fmt"
	"time"
)

func getOsName() string{
	return "mac"
}

func main() {
	// osをswitch内のみで使用する記述
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac")
	case "Windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Default")
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 24:
		fmt.Println("Night")

	}
}




