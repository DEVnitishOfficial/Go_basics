// loops and conditional

package main

import "fmt"

func main() {
	// age := 12

	// if age>=18 {
	// 	fmt.Println("you are adult");
	// }else if age >=13{
	// 	fmt.Println("You are teen ager")
	// }else{
	// 	fmt.Println("you are child")
	// }

	// for loop
	// for i := 1; i<=10; i++ {
	// 	fmt.Println(i)
	// }

	// infinit loop
	// for{
	// 	fmt.Println("Printing forever");
	// 	break;
	// }

	// looping on array/slices
	// nums := []int{10,20,30,40,50}

	// for index, value := range nums{
	// 	fmt.Println("index",index, "Value", value);
	// }

	marks := 75

	if marks > 90 {
		fmt.Println("A+")
	} else if marks > 80 {
		fmt.Println("A")
	} else if marks > 60 {
		fmt.Printf("B")
	} else {
		fmt.Printf("Fail")
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fruits := []string{"apple", "banana", "mango"}

	for index, value := range fruits {
		fmt.Println("Index", index, "value", value)
	}

	for i := range 3 {
		fmt.Println("printing range value", i)
	}
	for i, char := range "nitish" {
		fmt.Println("printing range value", i, char)
	}
}
