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
