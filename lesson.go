package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	// 新しいval
	m["banana"] = 300
	fmt.Println(m)
	// 新しいkey_val
	m["new"] = 500
	fmt.Println(m)

	// 存在しないkey
	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)
	//v := m["apple"]
	//fmt.Println(v)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 :=make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// nilのmap, memoryにmapがない
	var m3 map[string]int
	m3["pc"] = 5000
	fmt.Println(m3)

	var s []int
	if s == nil{
		fmt.Println("Nil")
	}



}
