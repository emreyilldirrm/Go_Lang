package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}
type Electronic struct {
	desi int
}
type Flower struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*3
}
func (flower *Flower) ShippingCost() int {
	return 5 + flower.desi*2
}
func main() {
	// book1 := &Book{desi: 10}
	// book2 := &Book{desi: 20}
	// fmt.Println(book1.ShippingCost())
	// fmt.Println(book2.ShippingCost())

	// books := []Book{{desi: 10}, {desi: 20}}
	// fmt.Println(calculateTotalShippingConstofBook(books))

	// electronic1 := &Electronic{desi: 10}
	// fmt.Println(electronic1.ShippingCost())

	// electronics := []Electronic{{desi: 10}, {desi: 20}}
	// fmt.Println(calculateTotalShippingConstofElectronic(electronics))

	//bir referansın birden farklı şekildeçalışması poliformizm
	// var product IShippable
	// product = &Book{desi: 10}
	// fmt.Println(product.ShippingCost())
	// product = &Electronic{desi: 10}
	// fmt.Println(product.ShippingCost())

	//farklı structları bir araya toplayabilir işlem yapabiliriz interface ile

	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Book{desi: 10},
		&Electronic{desi: 20},
		&Flower{desi: 10},
	}
	fmt.Println(calculateTotalShippingConst(products))

	//Video 2:30:28
}
func calculateTotalShippingConst(products []IShippable) int {
	total := 0

	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}

// func calculateTotalShippingConstofBook(books []Book) int {
// 	total := 0

// 	for _, book := range books {
// 		total += book.ShippingCost()
// 	}
// 	return total
// }
// func calculateTotalShippingConstofElectronic(electronics []Electronic) int {
// 	total := 0

// 	for _, electronic := range electronics {
// 		total += electronic.ShippingCost()
// 	}
// 	return total
// }
