package main

import "fmt"

func main() {

	var age int = 22
	var name string = "nitish kumar"
	height := 5.5
	var subject string = "PHYSICS, CHEMISTERY, MATH"
	var city string = "mumbai"
	var pincode int = 400001

	fmt.Println("name", name)
	fmt.Println("age : ", age)
	fmt.Println("height : ", height)
	fmt.Println("subject : ", subject)
	fmt.Printf("I live in, %s with pincode %d", city, pincode)

	main2()
}

func main2() {
	const myName = "nitish" // const is similr to js can't reassing the value also can't use shorthand property for it.

	var age int = 25
	var age2 int
	name := "Nitesh"
	height := 5.9
	age2 = 300

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("age2:", age2)
}
