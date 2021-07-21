package main

import 	"fmt"

type Book struct {
	title     string
	author    string
	publisher string
	isbnCode  string
	price     int
}

type DVD struct {
	title       string
	directer    string
	distributor string
	janCode     string
	price       int
}

func (b Book) display() {
	fmt.Printf("書名 = %s\n", b.title)
	fmt.Printf("著者 = %s\n", b.author)
	fmt.Printf("出版社 = %s\n", b.publisher)
	fmt.Printf("ISBN = %s\n", b.isbnCode)
	fmt.Printf("定価 = %d 円\n", b.price)
}
func (d DVD) display() {
	fmt.Printf("題名 = %s\n", d.title)
	fmt.Printf("監督 = %s\n", d.directer)
	fmt.Printf("販売元 = %s\n", d.distributor)
	fmt.Printf("JAN = %s\n", d.janCode)
	fmt.Printf("定価 = %d 円\n", d.price)
}

func main() {
	book1 := Book{"吾輩は猫である", "夏目漱石", "集英社", "4087520471", 1200}
	book2 := Book{"ノルウェーの森", "村上春樹", "講談社", "4062748681", 1800}
	dvd1 := DVD{"永遠の０", "山崎 貴", "アミューズメントソフトエンタテインメント", "4527427657830", 2900}
	dvd2 := DVD{"ブレードランナー　ファイナル・カット", "リドリー・スコット", "ワーナー・ブラザース・ホームエンターテイメント", "4548967342147", 950}
	fmt.Println()
	book1.display()
	fmt.Println()
	book2.display()
	fmt.Println()
	dvd1.display()
	fmt.Println()
	dvd2.display()
}
