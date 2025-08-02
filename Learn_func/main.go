package main

import "fmt"

func main() {

	greetUser("nitish")

	result := add(22, 53)
	fmt.Println(result)

	q, r := divide(47, 5)
	fmt.Println("quotent", q, "remainder", r)

	res := multiply(4, 70)
	fmt.Println(res)
}

func greetUser(name string) {
	fmt.Println("Hello,", name)
}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func multiply(a int, b int) (result int) {
	result = a * b
	return
}
