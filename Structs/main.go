// A struct in Go is a composite data type that groups together related values (fields) under a single name.

// it's similir to js object where we can put different typs of value in object.

// in js

/*
const user = {
  name: "Nitesh", // string
  age: 25 // numbers or int
}
*/

// in go
/*
type User struct {
    Name string
    Age  int
}

*/

package main

import "fmt"

type User struct {
	Name     string
	Age      int
	Email    string
	location string
}

func main() {

	user1 := User{
		Name:     "nitish",
		Age:      22,
		Email:    "nitish123@gmail.com",
		location: "noida",
	}

	fmt.Println(user1)    // read struts
	user1.Name = "kargil" // modifying value or write
	fmt.Println(user1)

	var user2 User
	fmt.Println("printing user2", user2.Age) // if not assign value we will got 0 in int and " " in string

	// calling greet method
	u := User{"Nitesh", 25, "nitesh@example.com", "delhi"}
	u.Greet()

	u1 := User{"Nitesh", 25, "nitesh@example.com", "delhi"}
	u1.updateEmail("updated@example.com")

	fmt.Println("priniting u1", u1)

}

// function with structs
func (u User) Greet() {
	fmt.Println("Hello,", u.Name)
}

//Pointer Receiver (to modify fields)
func (u *User) updateEmail(newEmail string) {
	u.Email = newEmail
}
