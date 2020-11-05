# Go-fintech
## 参考になるサイト
- https://godoc.org/
- https://ashitani.jp/golangtips/index.html
- http://www.openspc2.org/reibun/Go/1.1.1/
## Sec-3
- go get golang.org/x/tools/cmd/godoc
- go doc fmt

- Golandの使い方
1. File>New>Project
2. New Project
    - Location: /Users/mashimohroki/go/src/awesomeProject
    - SDK :1.14
3. awesomeProject > New(右クリック) > lesson.go
4. 右上 > Add Configuration(Edit Configuration) > Go Build
    - name: Go
    - Run Kind :Directory
  
    => apply
   
 ## Sec-6 
 ### Tips

- Go runタブの移動
    > ⚙ > Move to > Left , Right ...
- Vim 
```
option + shift + △ or ▽ : 行の入れ替え
p: ペースト
y: コピー(一文字)
yy: 1行コピー
vim [ファイル名]	起動する（ファイルを指定して、起動する）
:q	終了する
:w [ファイル名]	上書きする（名前をつけて保存する）
:q!	保存せずに終了する
h	左に移動する
j	下に移動する
k	上に移動する
l	右に移動する
w	前方に単語１つ分移動する
b	後方に単語１つ分移動する
W	スペース区切りで前方に単語１つぶん移動する
B	スペース区切りで後方に単語１つぶん移動する
O	行頭に移動する
$	行末に移動する
eg	１行目に移動する
G	最後の行に移動する
d <カーソル移動コマンド>	デリート（切り取りをする）
dd	行をデリート
y <カーソル移動コマンド>	ヤンク（コピーをする）
yy 行をヤンク
p	ペースト（貼り付ける）
i	カーソル位置の左に文字を入力する
※インサートモードに切り替える
a	カーソル位置の右に文字を入力する
※インサートモードに切り替える
Esc	ノーマルモードに切り替える
x	カーソル位置の文字を削除する
J	行を連結する
u	アンドゥ（元に戻す）
Ctrl + r	リドゥ（やり直す）
/ <文字列>	下方向に検索する
? <文字列>	上方向に検索する
n	次の検索結果に移動する
N	前の検索結果に移動する
:%s/<置換元文字列>/<置換後文字列>/g	置換する
vimtutor	チュートリアル
:help	ヘルプ
```
-  Run go : Ctrl + R

- GoFile structure
```
package project

import "library"

func name(){
    // Do something
}

func main(){
    name()
}
```

 ## Sec-7 
  ### Tips
- package
os/user: osライブラリ以下のuserライブラリ
 vim
 ```
x: 一文字削除
dw: 一単語削除
dd: 1行削除
D: 行末で削除
(num) + dd: カーソルから指定行数分削除
```


 ## Sec-8 
  ### Tips
 -  変数は複数宣言できる
```
var (
    var xf32 float32 = 1.2 //型を明示する
	i int = 1
	f64 float64 = 1.2
	s string = "test"
	t, f bool = true, false
)
```

 
 ## Sec-9 
  ### Tips
定数の性質 
```
// varは宣言時にコンパイラでオーバーフローする
//var big int =	9223372036854775807 + 1
// consはコンパイル実行時に宣言はオーバーフローしない
const big =	9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
```
 
## Sec-10 
  ### Tips

コードの整形
```
go fmt file.go
go fmt -w file.go
```

Goの変数ルール
一番長い変数に揃えた宣言をする

```
var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
```

## Sec-10 
  ### Tips
```
2yy + p:その位置から2行コピー+ペースト
yyp + .: 1行全体コピー&ペースト + 繰り返し
yyp + 10 +. : 1行全体コピー&ペースト + 10回繰り返しペースト

Shift+v で行単位のビジュアルモードに入り、範囲を指定する
範囲指定が終わったら、 y でコピーします。
ペーストしたい位置に移動して p で貼り付けます。

基本的にはpは次の行へのペーストになる
```
https://kaworu.jpn.org/vim/vim%E3%81%A7%E8%A1%8C%E3%82%92%E8%A4%87%E8%A3%BD%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95

string型の特徴

```
 	fmt.Println("Hello World"[0])         //ASCII表示
	fmt.Println(string("Hello World"[0])) //typecasting
// エスケープ文字列
	fmt.Println("\"")
	fmt.Println(`"`)
```

## Sec-13 
  ### Tips
GoのString<=>intの変換  
```

```

関数を右クリック> Go to > Declaration or Usage

```
func Atoi(s string) (int, error) {
	const fnAtoi = "Atoi"

	sLen := len(s)
	if intSize == 32 && (0 < sLen && sLen < 10) ||
		intSize == 64 && (0 < sLen && sLen < 19) {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0, &NumError{fnAtoi, s0, ErrSyntax}
			}
		}

		n := 0
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0, &NumError{fnAtoi, s0, ErrSyntax}
			}
			n = n*10 + int(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n, nil
	}
```

## Sec-14 
  ### Tips
配列とスライス

```
配列([num])はサイズ変更できない
    var b [2]int = [2]int{100, 200}
	b = append(b ,300) // appendできない
	fmt.Println(b)

slice([])はサイズ変更できる
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)

```

## Sec-16 
  ### Tips
スライスのmakeについて

Creating a slice with make
スライスは、組み込みの make 関数を使用して作成することができます。 これは、動的サイズの配列を作成する方法です。

make 関数はゼロ化された配列を割り当て、その配列を指すスライスを返します。

```
a := make([]int, 5)  // len(a)=5
make の3番目の引数に、スライスの容量( capacity )を指定できます。 cap(b) で、スライスの容量を返します:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4


// len = 3, cap = 3
a := make([]int, 3)
fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

//b: 0のスライスをメモリ確保, c: nill
b := make([]int, 0)
var c []int

//1: c = make([]int, 5)
//2: c = make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

// 1の結果
[0 0 0 0 0 0]
[0 0 0 0 0 0 1]
[0 0 0 0 0 0 1 2]
[0 0 0 0 0 0 1 2 3]
[0 0 0 0 0 0 1 2 3 4]
[0 0 0 0 0 0 1 2 3 4]

// 2の結果
[0]
[0 1]
[0 1 2]
[0 1 2 3]
[0 1 2 3 4]
[0 1 2 3 4]


```

書式指定子

```

// %d 基数10
    d := 5
    fmt.Printf("基数10 = %d\n", d)

// %T 値の型7
    fmt.Printf("int = %T\n", n)
    int = int

//
```


## Sec-17 
  ### Tips

```
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
=> 
map[apple:100 banana:200]
100

	// 新しいval
	m["banana"] = 300
	fmt.Println(m)
=>
map[apple:100 banana:300]

	// 新しいkey_val
	m["new"] = 500
	fmt.Println(m)
=>
map[apple:100 banana:300 new:500]

	// 存在しないkey
	fmt.Println(m["nothing"])
=>
0

	v, ok := m["apple"]
	fmt.Println(v, ok)
=>
100 true
	//v := m["apple"]
	//fmt.Println(v)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)
=>
0 false

	m2 :=make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)
=>
map[pc:5000]

	// nilのmap, memoryにmapがない
	//var m3 map[string]int
	//m3["pc"] = 5000
	//fmt.Println(m3)
=>
panic: assignment to entry in nil map

	var s []int
	if s == nil{
		fmt.Println("Nil")
	}
=>
Nil

```

## Sec-20
  ### Tips
  ブレークポイント: 赤丸
 デバッガー: 虫マーク 
 step over: ブレークポイントの次のステップに進む
 step into:　関数の中に入る。クリックするごとに1行ごと実行
 step out :
 
## Sec-26
  ### Tips
rangeで二つのパラメータを取り出すときの方法
  
```
//i: index, v: value
-       for i, v := range l{
-               fmt.Println(i, v)
-       }
-// valueを取り出すとき
-       for _, v := range l{
-               fmt.Println(v)
-       }
 m := map[string]int{"apple": 100, "banana": 200}
-// key ,valueを取り出す
-       for k, v := range m{
-               fmt.Println(k, v)
-       }
-// keyのみ
-       for k := range m{
-               fmt.Println(k)
-       }
-// valueのみ
-       for _, v := range m{
-               fmt.Println(v)

```

 
## Sec-27
  ### Tips
  
switch内のみで変数を使うとき
```
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
}

```

## Sec-28
  ### Tips

deferの実行順
```
func foo(){
 defer fmt.Println("world foo")

 fmt.Println("hello foo")
 
}


func main() {
 defer fmt.Println("world")

 foo()

 fmt.Println("hello")
}

//　出力
hello foo
world foo
hello
world

 fmt.Println("run")
 defer fmt.Println("1")
 defer fmt.Println("2")
 defer fmt.Println("3")
 fmt.Println("success")

//出力
run
success
3
2
1

 file, _ := os.Open("./lesson.go")
 // 一連の処理の締めを忘れずに書くことができる
 defer file.Close()
 data := make([]byte, 100)
 file.Read(data)
 fmt.Println(string(data))
```


## Sec-29
  ### Tips
```
log.Println("logging!")
log.Printf("%T %v", "test", "test")
 
 //Fatalは実行されるとそこでコードが終了する
log.Fatalln("error!")
fmt.Println("ok!")

//　出力
2020/09/23 22:51:44 logging!
2020/09/23 22:51:44 string test
2020/09/23 22:51:44 error!

// os.Openのエラーを出力してみる(エラーハンドリング)
//file, error :=os.Open("hogehogheohg") 
 _, err :=os.Open("hogehogheohg") //fileを変数として使わないときは"_"とする
 if err != nil{
  log.Fatalln("exit", err)
 }

// errorを出力 
2020/09/23 22:56:40 exit open hogehogheohg: no such file or directory

//logfileの出力
func loggingSettings(logFile string){
  logfile, _:= os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  multiLogFile := io.MultiWriter(os.Stdout, logfile)
  log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
  log.SetOutput(multiLogFile)
}

```


## Sec-30
  ### Tips
  
errorハンドリング

```
// errorがnilかどうかで処理を分ける
func Chdir(dir string) error {
	if e := syscall.Chdir(dir); e != nil {
		testlog.Open(dir) // observe likely non-existent directory
		return &PathError{"chdir", dir, e}
	}
	if log := testlog.Logger(); log != nil {
		wd, err := Getwd()
		if err == nil {
			log.Chdir(wd)
		}
	}
	return nil
}

// 100バイト分の配列を作成
	data := make([]byte, 100)
	// 100バイト分のlesson.goの中身を読み込む
	count, err := file.Read(data)
// Read => GoTo > Declaration
func (f *File) Read(b []byte) (n int, err error) {
	if err := f.checkValid("read"); err != nil {
		return 0, err
	}
	n, e := f.read(b)
	return n, f.wrapErr("read", e)
}

// errは コードの中で何度も :=　で初期化される
file, err := os.Open("./lesson.go")
  ~~~~~~	
count, err := file.Read(data) //fileのerrのため:=で初期化して上書きできる

err = os.Chdir("test") //ここはosに対する変数errのため初期化できない

```

## Sec-31
  ### Tips

panicとrecover
基本的にはこれらを使わないコード実装を推奨

```
panic("Unable to connect database!")

defer func() {
		s := recover()
		fmt.Println(s)
	}()
```

## Sec-32
  ### Tips
 
Exercise
配列のsortについて
http://midori5.hatenablog.com/entry/2017/06/11/194519
```

package main

import (
	"fmt"
	"sort"
)

func main() {
	l :=[]int{100, 300, 23, 11, 2, 4, 6, 4}
	sort.Ints(l)
	fmt.Println(l[0])

	m := map[string]int{
		"apple": 200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi": 90,
	}

	total := 0
	for _, v := range m {
		//fmt.Println(v)
		total = total + v
	}

fmt.Println(total)
}

```


## Sec-34
  ### Tips
ポインタ

メモリについて
https://murci.net/archives/319

図解

<img width="827" alt="スクリーンショット 2020-09-24 23 36 11" src="https://user-images.githubusercontent.com/54907440/94161742-f83acb00-fec0-11ea-963d-9531de0b4058.png">

```
var n int = 100
// アドレスと値の関係

fmt.Println(n)     //nの値
fmt.Println(&n)    //nのアドレス
fmt.Println(*&n)   //nのアドレスに格納された値
fmt.Println(&*&n)  //nのアドレスに格納された値のアドレス


```

## Sec-35
  ### Tips
makeとnewについて
- make: ポインタを返さない
- new : ポインタを返す

```

s := make([]int, 0)
	fmt.Printf("%T\n", s) // []int


	m := make(map[string]int)
	fmt.Printf("%T\n", m) // map[string]int


	ch := make(chan int)
	fmt.Printf("%T\n", ch) // chan int

	var p *int = new(int)  
	fmt.Printf("%T\n", p)  // *int

	var st = new(struct{})
	fmt.Printf("%T\n", st) // *struct {}

// メモリにアドレス領域のみ確保したいとき　
var p *int = new(int)
	fmt.Println(p)

// p2はnil
var p2 *int
	fmt.Println(p2)
```

## Sec-36
  ### Tips
structについて
Objectのようなもの
 
 ```
// パラメータと型を宣言する
type Vertex struct {
	X int
	Y int
	S string
}

v := Vertex{X:1, Y:2}
	fmt.Println(v.X, v.Y) //1 2
	v.X = 100
	fmt.Println(v.X)      // 100 

	v2 := Vertex{X:1} 
	fmt.Println(v2) // {1 0 }


	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3) // {1 2 test} 

    v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4) // main.Vertex {0 0 }

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5) // main.Vertex {0 0 }

	v6 := new(Vertex) //ポインタ
	fmt.Printf("%T %v\n", v6, v6) // *main.Vertex &{0 0 }

    v7 := &Vertex{}  // newよりもポインタを明示的にする
	fmt.Printf("%T %v\n", v7, v7) // *main.Vertex &{0 0 }
}


func changeVertex(v Vertex)  {
	v.X = 1000
}

// ポインタでの宣言時の違い
func changeVertex2(v *Vertex)  {
	//(*v).X = 1000
	v.X = 1000 // structはvを*vとして自動的に変換
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v) // {1 2 test}

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2) // &{1000 2 test}
    fmt.Println(*v2) // {1000 2 test}}
```


## Sec-39
  ### Tips
メゾッド

```
type Vertex struct {
	X, Y int
}

// Object-Orientedライクなメゾッド宣言
func (v Vertex) Area()  int{
	return v.X * v.Y
}



//　関数
func Area(v Vertex)  int {
	return v.X * v.Y
}

func main() {
	 
    v := Vertex{3, 4}
	v2 := Vertex{5, 6}
	fmt.Println(Area(v))
	fmt.Println(v.Area()) //メゾッド呼び出し 
	fmt.Println(v2.Area())
}
```

・ポインタレシーバ

```
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}


func main() {
    v := Vertex{3, 4}
	fmt.Println(v.Area()) //12
	v.Scale(10)
	fmt.Println(v.Area())
}

```

・値レシーバ

```
func (v Vertex) Area()  int{
	return v.X * v.Y
}

func main() {
    v := Vertex{3, 4}
	fmt.Println(v.Area()) //12
	v.Scale(10)           //アドレスに格納されている値を上書き
	fmt.Println(v.Area()) //1200

```


## Sec-39
  ### Tips
Vimコマンド 
- 置換
```
:%s/対象文字列/置換する文字列/g
=> enter
```
 
コンストラクタ
golang document :https://golang.org/pkg/math/rand/#New

```
func New

func New(src Source) *Rand

New returns a new Rand that uses random values from src to generate other random values.
```

```
func New(x, y int) *Vertex{
	return &Vertex{x, y}
}

func main() {
	//v := Vertex{3, 4}
	v := New(3, 4) // 
	v.Scale(10)
	fmt.Println(v.Area())
}
```

## Sec-40
  ### Tips
Embedded
構造体の埋め込み
クラスの継承のようなもの

参考記事: https://qiita.com/k-penguin-sato/items/62dfe0f93f56e4bf9157#structs%E6%A7%8B%E9%80%A0%E4%BD%93

```

// Vertexを埋め込み(embedded)
type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D()  int{
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D{
	// Vertexを宣言
    return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())
}
```

## Sec-41
  ### Tips
non structメゾッド
```

type MyInt int

func (i MyInt) Double() int{
	fmt.Printf("%T %v\n", i, i) // main.MyInt 10

	fmt.Printf("%T %v\n", 1, 1) // int 1

	return int(i * 2) // intへのキャスティングが必要
} 

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double()) //20
}
```


## Sec-43
  ### Tips
 インターフェース
 あるクラス(構造体)にメゾッドを実装することを保証する
 "ふるまいの共通化"
 参考記事: https://www.gixo.jp/blog/5159/
 
 interfaceで定義したメソッドをクラス(Goでは構造体)内で必ず使用(オーバーライド)しなければならない
 

```
// メゾッドのひな型のようなもの
type Human interface {
	Say()
}

type Person struct {
	Name string
}


// PersonはSay()をもつ
func (p Person) Say(){
	fmt.Println(p.Name)
}

// Person
func (p *Person) Say(){ // Say()を宣言したので、Humanインターフェースを実装したことになる
    p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
}

// interfaceで指定したメゾッドが必要になる
func main() {
	var mike Human = Person{"Mike"}
	mike.Say()

	//var mike Human = &Person{"Mike"} //アドレスの値として受け取る  
	//mike.Say()
}

// 構造体Dog
type Dog struct {
	Name string
}


func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human){
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	var dog Dog = &Dog{"hachi"}
	DriveCar(mike)
	DriveCar(x)

    // DogインスタンスはHumanインターフェースを持たないのでDriveCarは実行できない
	DriveCar(dog)   //Dog does not implement Human (missing Say method)
	//var mike Human = &Person{"Mike"}
	//mike.Say()
```

ダックタイピング:
Go言語にはインタフェースがある。
明示的に実装しなくてもインタフェースが定義するメソッドをすべて実装していれば、
そのインタフェースを実装していることになる（ダックタイピング）。
```
package main

import "fmt"

type Person struct {
    FirstName string
    LastName  string
}

func (p *Person) Name() string {
    return p.FirstName + " " + p.LastName
}

type Named interface {
    Name() string
}

func printName(named Named) {
    fmt.Println(named.Name())
}

func main() {
    person := &Person{"Tarou", "Yamada"}
    printName(person)
}
```
https://qiita.com/tenntenn/items/e04441a40aeb9c31dbaf

"静的型付け言語であるJavaやC#の概念で例えると、
オブジェクトがあるインタフェースのすべてのメソッドを持っているならば、
たとえそのクラスがそのインタフェースを宣言的に実装していなくとも、
オブジェクトはそのインタフェースを実行時に実装しているとみなせる、
ということである。"
参考:
https://www.weblio.jp/content/%E3%83%80%E3%83%83%E3%82%AF%E3%82%BF%E3%82%A4%E3%83%94%E3%83%B3%E3%82%B0 



## Sec-44
  ### Tips
タイプアサーション
インターフェースをintやstringに変換する

https://go-tour-jp.appspot.com/methods/15
"型アサーション は、インターフェースの値の基になる具体的な値を利用する手段を提供します。

```
t := i.(T)
```
この文は、インターフェースの値 i が具体的な型 T を保持し、基になる T の値を変数 t に代入することを主張します。

i が T を保持していない場合、この文は panic を引き起こします。

インターフェースの値が特定の型を保持しているかどうかを テスト するために、型アサーションは2つの値(基になる値とアサーションが成功したかどうかを報告するブール値)を返すことができます。

```
t, ok := i.(T)
```

i が T を保持していれば、 t は基になる値になり、 ok は真(true)になります。

そうでなければ、 ok は偽(false)になり、 t は型 T のゼロ値になり panic は起きません。

この構文と map から読み取る構文との類似点に注意してください。"

```
func do(i interface{}){
	ii := i.(int) //interface型をint型へ変換
	ii *=  2
	fmt.Println(ii)
}

func main() {
	var i interface{} = 10 // interface型
	do(i)
}
=> 20
```

文字列の場合

```
func do(i interface{}){
	ss := i.(string)
	fmt.Println(ss + "!")
}

func main() {
	do("Mike")
}
=> Mike!
``````

swich type文

```
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}


func main() {
	do(10)
	do("Mike")
	do(true)
}

=>
20
Mike!
I don't know bool

```

Type Conversion(キャスト)

```
var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
```


## Sec-45
  ### Tips
Stringer
>Stringer インタフェースは、stringとして表現することができる型です。 fmt パッケージ(と、多くのパッケージ)では、変数を文字列で出力するためにこのインタフェースがあることを確認します。

出力の結果がString()で定義された
文字列になる。
```
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

=> Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)


type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main()  {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}

=> "My name is Mike.
```
## Sec-46
  ### Tips
カスタムエラー

カスタムエラーの引数のタイプには*をつけて、
出力する際には&をつける。
=> エラーの内容に違いが出てしまうため。

```
type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string{
	/*return e.Username*/
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	// Something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}

=> User not found: mike
```

## Sec-47
演習
```
ex1 7を出力するメゾッド

package main

import (
	"fmt"
)

type Vertex struct{
	X, Y int
}

func  (v Vertex) Plus() int {
	return v.X + v.Y
}

func main(){
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
}

ex2) X is 3! Y is 4! と表示されるStringer

package main

import (
	"fmt"
)

type Vertex struct{
	X, Y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!", v.X, v.Y)
}
func main(){
	v := Vertex{3, 4}
	fmt.Println(v)
}
```

## Sec-49
### Tips
goroutine
軽量のスレッド、並列処理のこと。
https://qiita.com/ruiu/items/dba58f7b03a9a2ffad65
> wg.Done()はdefer wg.Done()として関数の先頭に書いてしまうのが良いとされています。
  deferは関数の呼び出しをリストで保存して、関数が終了したときに順次呼び出します。
```
package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup)  { //　sync.WaitGroupは複数のgoroutineの完了を待つための値
        defer wg.Done()
	for i := 0; i < 5; i++{
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
    // wg.Done()	
}

func normal(s string){
	for i := 0; i < 5; i++{
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main()  {
	var wg sync.WaitGroup
	wg.Add(1)                  // wgに一つの並列処理があることを伝える           
	go goroutine("world", &wg) // go:並列処理の宣言 何もしないと処理が終了しなくてもプログラムが終わることがある
	normal("hello")
	//time.Sleep(2000 * time.Millisecond)
	wg.Wait()                  // wgがDone()するまで待つ
}
```
## Sec-50
### Tips
channel

mainとgoroutineの間でデータを
やり取りするための仕組み
![channel](https://user-images.githubusercontent.com/54907440/98129473-806ab200-1efc-11eb-8fb4-98ede63895b1.png)

2つのgoroutineに１つのチャネル
2つのチャネルを使うこともできる。
(一つはint, もう一つはstringなど)
![channel2](https://user-images.githubusercontent.com/54907440/98131055-43072400-1efe-11eb-8a27-0b635f7a8a57.png)

## Sec-51
### Tips
Buffered Channels
rangeでchannelを呼び出すときは、
closeでchannelを閉じないとエラーが出る
```
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
```
## Sec-52
### Tips
channelのrangeとclose
バッファありとバッファなしチャネルの違い
https://program.sakaiboz.com/golang/goroutine/unbuffered-channel-and-buffered-channel/
```
func goroutine1(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
		c <- sum
	}
	close(c) //レンジに値がなくなった時にチャネルを閉じる、閉じないとエラー
}

func main()  {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	for i := range c{
		fmt.Println(i)
	}
}
```
![range_and_close](https://user-images.githubusercontent.com/54907440/98254492-b2911800-1fbf-11eb-8e82-c357b997e65f.png)


## Sec-53
### Tips
producerとconsumer
producerのgoroutineの処理結果を
consumerのgoroutineで処理するというイメージ。
![producer_consumer](https://user-images.githubusercontent.com/54907440/98256359-d0f81300-1fc1-11eb-93b3-cb12288dd935.png)
2つのgoroutineを用意している
```
func producer(ch chan int, i int)  {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup)  {
	for i := range ch{
		fmt.Println("process", i * 1000)
		wg.Done()
	}
	fmt.Println("#############")
}

func main()  {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

```

