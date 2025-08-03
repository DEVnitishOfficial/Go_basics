
# Goal : Learning structs

* In other programming languages, we have classes and objects like java, javascript, python, c++ in go, we have structs which are similar to classes, which help us to create custom data types
// where we group different data types together like we do in javascript objects

* A struct in Go is a composite data type that groups together related values (fields) under a single name.

**Define structs**
```go
type Product struct {
	name    string
	price   int
	company string
}
```
**creating product from structs** 
```go
product1 := Product{
		name:    "I phone 15 pro",
		price:   100500,
		company: "apple",
	}
```

* In other programming language you have the consturctor or some default constructor but in go there is no constructor function available in go programming language. if we have to create an object we can directly create object from the sturcts whatever we have defined.


* But if we have to define some rules in constructor then we can't do it in go, but here is a catch, that we can define a function in such a way which will work as a contructor. like below one 

// this fucntion will take some paramerter and out of that we will create a new object
```go
func newProduct(name string, price int, company string) Product {
	// creating object
	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	// now i want to return the above particular object, how to do?
	// define the type of return type ----> Procuct and just return p
	return p // this return the copy of the product
}
```


# Goal ======>>>> Understanding in GoLang everthing is passed as copy not reference


* Consider the below example and observe that how it always return the copy either you pass any object as a parameter to the function or returning any value from any function, it always give you the copy not the reference.

* Below is the explanation of the code and answer of how it's happening?.

```go
package main

import "fmt"

type Product struct {
	name    string
	price   int
	company string
}

// this fucntion will take some paramerter and out of that we will create a new object
func newProduct(name string, price int, company string) Product {
	// creating object
	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	// now i want to return the above particular object, how to do?
	// define the type of return type ----> Procuct and just return p
	return p  // this is returning the copy of product p

}

func fun(p Product) {
	// here p is the copy of product "p" which is passed to this fun function
	p.name = "mackBook pro"

	fmt.Println("name inside fun function", p.name)
}

func main() {

	// creating product from Product structs

	// here in struct we dont have the new keyword so simply using the below syntax we can create a new product
	product1 := Product{
		name:    "I phone 15 pro",
		price:   100500,
		company: "apple",
	}

	fmt.Println("product name before fun function call>>>>>>", product1.name)

	fun(product1)

	fmt.Println("product name after fun function call>>>>>>", product1.name)

}
```

**Explanation** 

* In the above code we are calling a function "fun" before and after the execution of "fun" function here we have created an object named "product1" from the Product struct and passed that object to the fun function.

* Inside the fun function we are receiving that object as "p" and override the name property from "I phone 15 pro" to "mackBook pro" but when we print the value of name property after function execution we see the same "I phone 15 pro" name which means "product1" is passed a value not reference.

* And the changes is reflected only inside the "fun" function scope, try to run and see the output.



# Next Goal : what if we wnat to pass references not copy of object ------>>>> Use pointers and addresses

* Our next goal is to learn how we can passs the same object not the copy, since by default Go return the copy of the object, using the pointers and addresses we can pass the same object like below one :

```go
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

func main() {

	// creating a new product using the constructor function
	// here new_product is a pointer so using . operator to access the fields and also we can use * operator to dereference the pointer like (*new_product).name
	// but it is not recommended to use * operator to dereference the pointer as it makes the code less readable
	// so we will use the . operator to access the fields of the product
	new_product := newProduct("I phone 15 pro", 100500, "apple")

	fmt.Println("product name before pass_by_reference function call>>>>>>", new_product.name)

	pass_by_reference(new_product)

	fmt.Println("product name after pass_by_reference function call>>>>>>", new_product.name)

}
```

* When you see *Product, it signifies a pointer to a Product struct.

* What is a pointer? A pointer is a variable that stores the memory address of another variable. Instead of holding the actual data (like the name, price, and company), it holds the location where that data is stored in your computer's memory.


* In the above code example we are returning the address of the created object and also passing the reference of the object inside the function so that, whatever changes of modification happening that should be inside the same object that's why you see the following changes : 

```go 
product name before pass_by_reference function call>>>>>> I phone 15 pro
name inside pass_by_reference function>>>>> samsung galaxy s24
product name after pass_by_reference function call>>>>>> samsung galaxy s24
```


