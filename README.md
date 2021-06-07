# Go

#Go 言語の特徴
**メリット**

- C や JAVA と構文が似ている
- 直接機械語に変換されるから処理速度が速い
- 処理が簡易的で複数処理を同時に行うことができる
- ライブラリが豊富
- 安全性が高い(メモリの開放忘れなどが発生しにくい)

**デメリット**

- Genrics が存在しない(引数の関数の型がきまっている)
- 例外処理がない
- コードの継承がない(オブジェクト指向言語ではなく関数型プログラム言語)

# ビルド実行コマンド

- $go build ファイル名.go → .\ファイル名(go build でできたファイル)
- $go run ファイル名.go(デバック時に使用)

# データ型

- 論理値型(bool) True か False どちらかの値を取る
- 整数型(int32(+) uint32(-)) 整数を保存するために使用
- 実数型(float64) 実数を保存する
- 複素数型(complex64) 実数と虚数を保存
- 文字列型(string) 0 文字以上の文字からなる文字列を表す(日本語の場合,1 文字 3B になる)
- 配列型 同じ型の要素を並べたもの
- スライス型 配列内の連続した領域への参照
- 構造体型 フィールドと呼ばれる要素の集まり
- ポインター型(\*int) メモリ上のオブジェクトを指し示す値
- 関数型 呼出可能な一連のプログラミングコード
- インターフェース型 実行されるコードがない関数
- マップ型 キーと値を持つ値
- チャンネル型 並列処理で使用
- 型の変換

# 変数と定数

- 変数定義 var 変数名 型 (文字の場合は配列として扱われる)
- 変数定義の省略 変数名 =: 値
- 定数の定義 const 定数名 = 値

# 演算子

- 単項演算子 ++や--は 1 つのステートメントとして扱う
  ```Go:ステートメント
  a++
  c = a // c = a++ はエラー
  ```

# コンソール入力

- fmt.Println() 出力して改行

```Go:fmt.Println
var name string = "sampleくん"
fmt.Println("こんにちは", name, "。")
```

- fmt.Printf() 出力する書式を指定

```GO:fmt.Printf
var r rune='ゆ'
fmt.Printf("%d %x %f %s", r, r, r, r)
```

- fmt.Scan() キーボードからの入力を変数に格納

```GO:fmt.Scan()
var name string
fmt.Printf("Name:")
fmt.Scan(&name)
fmt.Printf("Hello,%s!", name)
```

- fmt.Scanf() 書式を指定してキーボードからの入力を変数に格納

```Go:fmt.Scanf()
~
fmt.Scanf("%d %d", &n, &m)
fmt.Printf("%dと%dの合計は%d\n", n, m, n+m)
```

- スキャナー 別ファイルからの入力を受け付ける
  scanner.go 参照

# 条件分岐

- if 文 if.go 参照
- switch 文 switch.go 参照
- 無条件分岐 無条件に指定したラベルの場所にジャンプ(使わないほうがいいらしい)

```GO:go to文
LOOP:
if n > 5{
  goto LOOPEND
}
```

# 繰り返し

- for 文 for.go 参照
- for ‥ range 文 range.go 参照

# 配列

- 配列の個数が足りない時、0 で補われる

```Go:配列
// 配列の宣言 var name [size]type
var a [5]int = [5]int{1, 3, 5, 7, 9}
// 配列の個数が未定の場合
b := [...]int{2, 4, 5, 6, 7}
```

# スライス

- 配列の 1 部又は全部を指し示す参照型データ 特殊な挙動は sliceprac.go 参照

```Go:スライス
// var name = []type{}
var c = a[0:3]
```

- make() スライスの作成(あんまり使わない)

```Go:make()
// make([]type, length, capacity)
var d = make([]int, 5, 20)
```

- cap() スライスの容量を調べる(保存できる要素数である容量)

```Go:cap()
fmt.Println(cap(a))
```

- len() サイズを調べる(保存されている要素数である要素数)

```Go:len()
fmt.Println(len(a))
```

- copy() スライスの要素を別のスライスにコピーする

```Go:copy()
copy(c, d)
```

# マップ
* キーと値のペアを複数保存する構造
```Go:マップ
// map [keytype] type
m := map[string]int{"sample": 100, "test": 200} // map[sample:100 test:200]
// キーの存在を調べる
v, ok := m["sample"]
```
* make() マップの作成(使わない)
```Go:make()
var m2 = make(map[string]int)
// 要素追加後自動的にソートされる
m2["sample"] = 20
m2["sample2"] = 21
m2["sample"] = 24
```
* delete() 指定したキーの要素を削除
```Go:delete()
delete(m, "sample")
```

# 構造体
* 複数の値を意味のあるまとまりとして新しい型を定義する事ができる
```Go:構造体
// 構造体を定義(先頭を大文字にすると他のパッケージからアクセス可能)
type User struct {
    // フィールド
    Name  string
    Score int
}
```

* new() newで構造体分の領域を確保して、初期化して、そのポインタを返す
```Go:new()
func main(){
  var u1 = new(User)
  // 構造体の値にデータをセット パターン1
  *u1 = User("sample", 100)
  // パターン2
  u1.Name = "sample2"
  u1.Score = 200
  // パターン3
  var u1 [3]User = [3]User{
    User{"sample3", 150},
    User{"sample4", 230},
    User["sample5", 50],
  } 
```

* 構造体をもつ構造体 JAVAの継承の様なもの embedding.goを参照

# 関数
* math,string,timeパッケージに関するもの func.go参考
* 関数の定義
```Go:関数の定義
func name(n int) int{
  return 2 * n
}
```
* 複数の値を返す関数
```Go:複数の値を返す関数
func name(n int) (int, int){
  
}
```

2021/06/07 00:44:09