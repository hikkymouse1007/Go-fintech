package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./lesson.go")
	if err != nil{
		log.Fatalln("Error!")
	}
	defer file.Close()
	// 100バイト分の配列を作成
	data := make([]byte, 100)
	// 100バイト分のlesson.goの中身を読み込む
	count, err := file.Read(data) //errを上書きする
	if err != nil{
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	err = os.Chdir("test") //ここはerrを初期化できない
	if err != nil{
		log.Fatalln("Error")
	}
}
