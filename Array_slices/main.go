package main

import "fmt"

func main() {

	var numArr [8]int

	numArr[0] = 30
	numArr[4] = 300
	numArr[7] = 20
	numArr[6] = 45
	fmt.Println("seemy array", numArr)

	var mystr [5]string

	mystr[0] = "beauty"
	mystr[2] = "nitish"
	mystr[4] = "kargil"
	mystr[1] = "amrit"
	fmt.Println(mystr)

	// shortcut
	nums := [3]int{1, 2, 3}

	str := [3]string{"hi", "hello", "this"}

	fmt.Println("myNum", nums)
	fmt.Println("stringvalues", str)

	num1 := []int{45, 46, 23, 58, 20, 50}

	sl := num1[1:4]
	fmt.Println("see my slice", sl)

	num1 = append(num1, 100, 500, 800, 259)
	fmt.Println("updatedNum", num1)

	// fruits := []string{"apple", "banana", "cherry"}

	// for index, fruit := range fruits {
	// 	fmt.Println("Index:", index, "Fruit:", fruit)
	// }

	// for index, val := range fruits {
	// 	fmt.Println("index", index, "fruits", val)
	// }

	fruits1 := []string{"f1", "f2", "f3", "f4", "f5", "f6"}

	fruits1 = append(fruits1[:2], fruits1[3:]...)
	fmt.Println("deleted f3 fruits", fruits1)

	for _, val := range fruits1 {
		fmt.Println("val", val)
	}

	// solving the practice problem
	// 1. **Array Basics:**
	// - Declare an array of 4 strings and print all names.
	// - Try modifying the second element and print the array.

	nameStr := [4]string{"beauty", "raman", "nitsh", "mohit"}

	nameStr[2] = "kargil"

	fmt.Println("printin names", nameStr)

}
