package main

import "fmt"

type printer interface {
	print()
}

type list []printer

type Game struct {
	name    string
	playing bool
	price   float32
}

type Book struct {
	name    string
	author  string
	reading bool
	price   float32
}

func (game *Game) print() {
	fmt.Println("[ name: ", game.name, " ]")
}

func (book *Book) print() {
	fmt.Println("[ name: ", book.name, ", author: ", book.author, " ]")
}

func (book *Book) discount() float32 {
	return book.price * 0.8
}

func (p list) discount1() {
	// only book has discount
	for _, v := range p {
		book, isBook := v.(*Book)
		if isBook {
			fmt.Println("New price: ", book.discount())
		}
	}
}

func (p list) discount2() {
	for _, v := range p {
		// check that v has method named discount() which returns float32
		item, ok := v.(interface{ discount() float32 })
		if ok {
			fmt.Println("New price: ", item.discount())
		}
	}
}

func (p list) discount3() {
	type discounter interface {
		discount() float32
	}
	for _, v := range p {
		// check that v satifies discounter interface
		item, ok := v.(discounter)
		if ok {
			fmt.Println("New price: ", item.discount())
		}
	}
}

func main() {
	book := Book{name: "Go", author: "dont know", reading: true, price: 100}
	game := Game{name: "go go", playing: false, price: 350}

	var p list
	p = append(p, &book)
	p = append(p, &game)

	for _, v := range p {
		v.print()
	}

	// type assertion
	p.discount1()
	p.discount2()
	p.discount3()
}
