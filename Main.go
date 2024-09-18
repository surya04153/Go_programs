package main

import "fmt"

var sessionName string = "Golang Session"

func main() {
	fmt.Println("hello world Go!!!!")
	var firstName string = "surya"
	var age int8 = -127
	fmt.Println(firstName, age, sessionName)
	sessionName = "mysession"
	fmt.Println(firstName, age, sessionName)
	fmt.Printf("Hi %v and %v \n", firstName, sessionName)
	fmt.Printf("Hi %T and %T \n", firstName, age)
	var b bool = true

	fmt.Printf("%b", age)
	fmt.Println("b", b)

	var x float32 = 3.4e+3
	fmt.Println(x)

	w := 3.4e+3
	fmt.Println(w)

}
