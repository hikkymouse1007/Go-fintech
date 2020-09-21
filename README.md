# Go-fintech
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
- 
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
```
Creating a slice with make
スライスは、組み込みの make 関数を使用して作成することができます。 これは、動的サイズの配列を作成する方法です。

make 関数はゼロ化された配列を割り当て、その配列を指すスライスを返します。

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







~~~~

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
