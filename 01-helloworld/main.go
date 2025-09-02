package main

import "fmt"

func main() {

	// string
	fmt.Print("Hello !!!!\n")
	fmt.Println("This will print in new line without \\n ")

	// string concat
	fmt.Print("This " + "is " + "go" + " Tut")

	// int, float
	fmt.Println(10)
	fmt.Println(10.99)
	fmt.Println(10 + 12)
	fmt.Println(10 + 12.33)

	// bool
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(false && true)
	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(true || false)
}
