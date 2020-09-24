package main

import "fmt"

func one(x *int){
	// 引数のアドレスのメモリの中身を1にする
	*x = 1
}

func main()  {
	var n int = 100
	one(&n)

	fmt.Println(n)
	fmt.Println(&n)
	fmt.Println(*&n)
	fmt.Println(&*&n)

	/*var n int = 100
	fmt.Println(n)

	// nのメモリアドレスを表示
	fmt.Println(&n)

	// nのメモリアドレスを代入
	var p *int = &n

	//　nのメモリアドレスを表示
	fmt.Println(p)

	//nのメモリアドレスに格納されている値を表示
	fmt.Println(*p)
	*/
}

