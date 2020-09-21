package main

import "fmt"

func add(x int, y int) (int, int) {
	//fmt.Println("add function")
	//fmt.Println(x + y)
	return x + y, x - y
}

// 複数の戻り値を返すときは変数を明示する(result int, tax int etc....)
func cal(price, item int) (result int) {
	result = price * item
	return result
}

//func cal(price, item int) (int) {
//	result := price * item
//	return result
//}

func convert(price int) float64{
	return float64(price)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	// 無名関数
	f := func(x int) {
		fmt.Println("inner func", 1)
	}
	f(1)

	// 関数の直接実行ができる
	func(x int){
		fmt.Println("inner func", x)
	}(1)
}
