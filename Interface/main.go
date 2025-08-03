package main

import "fmt"

// sellableProduct is an interface that defines the methods that a product struct must implement
// if the struct not implement all these mehtods that are defined inside the interface then it will throw an error
type sellableProduct interface {
	buy()
	getDiscount() int
}

type Product struct {
	name    string
	price   int
	company string
}

func newProduct(name string, price int, company string) *Product {
	// creating object
	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	return &p
}

// below two methods buy() and getDiscount() methods are part of Product struct
// so we need to implement these methods to satisfy the sellableProduct interface

// implementing the buy method of the sellableProduct interface
func (p *Product) buy() {
	fmt.Println("buying product:", p.name)
}

// implementing the getDiscount method of the sellableProduct interface
func (p *Product) getDiscount() int {
	discount := p.price * 20 / 100
	fmt.Println("getting discount for product:", p.name)
	return discount
}

func check_discount_and_buy(p sellableProduct) {
	discount := p.getDiscount()
	if discount > 30 {
		fmt.Println("discount is good enough, buying the product")
		p.buy()
	} else {
		fmt.Println("discount is not good enough, not buying the product")
		return
	}
}

func main() {

	// creating a new product using the constructor function
	new_product := newProduct("I phone 15 pro", 100500, "apple")

	check_discount_and_buy(new_product)

	fmt.Println(new_product.name)
}
