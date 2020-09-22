package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++{
		fmt.Println(i, l[i])
	}
//i: index, v: value
	for i, v := range l{
		fmt.Println(i, v)
	}
// valueを取り出すとき
	for _, v := range l{
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
// key ,valueを取り出す
	for k, v := range m{
		fmt.Println(k, v)
	}
// keyのみ
	for k := range m{
		fmt.Println(k)
	}
// valueのみ
	for _, v := range m{
		fmt.Println(v)
	}
}




