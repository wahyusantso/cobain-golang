package main

import "fmt"

func main() {
	name := "John Doe"

	if name == "John" {
		fmt.Println("Hello, John!")
	} else {
		fmt.Println("Go Ahead!")
	}

	if length := len(name); length > 5 { //if short statement untuk menghitung panjang string name
		fmt.Println("Name is too long")
	}
}
