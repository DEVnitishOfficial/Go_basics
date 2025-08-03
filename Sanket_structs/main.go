package main

import "fmt"

type Product struct {
	name    string
	price   int
	company string
}

// newProduct is a constructor function to create a new Product object
func newProduct(name string, price int, company string) *Product {
	// creating object
	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	return &p

}

func pass_by_reference(p *Product) {
	// here p is the reference of product "p" which is passed to this pass_by_reference function
	p.name = "samsung galaxy s24"
	fmt.Println("name inside pass_by_reference function>>>>>", p.name)
}

// define the member function for Product struct
// here p refers to new_product object
// Here display is not a normal function, it is a member function of the Product struct
func (p *Product) display() {
	fmt.Println("product Details:")
	fmt.Println("Name:", p.name)
	fmt.Println("Price:", p.price)
	fmt.Println("Company:", p.company)
}

func main() {

	// creating a new product using the constructor function
	// here new_product is a pointer so using . operator to access the fields and also we can use * operator to dereference the pointer like (*new_product).name
	// but it is not recommended to use * operator to dereference the pointer as it makes the code less readable
	// so we will use the . operator to access the fields of the product
	new_product := newProduct("I phone 15 pro", 100500, "apple")

	fmt.Println("product name before pass_by_reference function call>>>>>>", new_product.name)

	pass_by_reference(new_product)

	fmt.Println("product name after pass_by_reference function call>>>>>>", new_product.name)

	new_product.display() // calling the member function to display product details

}
