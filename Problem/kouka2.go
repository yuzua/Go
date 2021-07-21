package main

import "fmt"

type Media struct {
	title string
	price int
}

type Book struct {
	Media
	auther string
	publisher string
	isbnCode string
}

type DVD struct {
	Media
	directer string
	distributor string
	janCode string
}

func (m Media) display() {
	fmt.Printf("タイトル = %s\n", m.title)
	fmt.Printf("定価 = %d 円\n", m.price)
}

func (b Book) display() {
	b.Media.display()
	fmt.Printf("著者 = %s\n", b.auther)
	fmt.Printf("出版社 = %s\n", b.publisher)
	fmt.Printf("ISBN = %s\n", b.isbnCode)
}

func (d DVD) display() {
	d.Media.display()
	fmt.Printf("監督 = %s\n", d.directer)
	fmt.Printf("販売元 = %s\n", d.distributor)
	fmt.Printf("JAN = %s\n", d.janCode)
}

func main() {
	book1 := Book{Media{"吾輩は猫である", 1200}, "夏目漱石", "集英社", "4087520471"}
	book2 := Book{Media{"ノルウェーの森", 1800}, "村上春樹", "講談社", "4062748681"}
	dvd1 := DVD{Media{"永遠の０", 2900}, "山崎 貴", "アミューズメントソフトエンタテインメント", "4527427657830"}
	dvd2 := DVD{Media{"ブレードランナー　ファイナル・カット", 950}, "リドリー・スコット", "ワーナー・ブラザース・ホームエンターテイメント", "4548967342147"}
	fmt.Println()
	book1.display()
	fmt.Println()
	book2.display()
	fmt.Println()
	dvd1.display()
	fmt.Println()
	dvd2.display()
}