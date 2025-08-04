package main

import (
	"fmt"

	app "GoModDemo/app"
)

func main() {

	// s := hello()
	// fmt.Println(s)

	c1 := app.NewComplex(3.0, 4.0)
	c2 := app.NewComplex(1.0, -2.0)

	// Display original numbers
	fmt.Println("C1:", c1.Display())
	fmt.Println("C2:", c2.Display())

	// adding two complex numbers
	sum := c1.Add(*c2)
	fmt.Println("C1 + C2 =", sum.Display())

	// subtracting two complex numbers
	diff := c1.Subtract(*c2)
	fmt.Println("C1 - C2 =", diff.Display())

	// multiplying two complex numbers
	product := c1.Multiply(*c2)
	fmt.Println("C1 * C2 =", product.Display())

	// dividing two complex numbers
	quotient, err := c1.Divide(*c2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("C1 / C2 =", quotient.Display())
	}

}
