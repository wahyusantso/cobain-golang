package main

import "fmt"

func main() {
	name := "John Doe"

	switch name {
	case "John":
		fmt.Println("Hello, John!")
	case "Doe":
		fmt.Println("Hello, Doe!")
	default:
		fmt.Println("Go Ahead!, john is not here")
	}

	switch length := len(name); length > 10 { //switch short statement untuk menghitung panjang string name
	case true:
		fmt.Println("Name is too long")
	default:
		fmt.Println("Name is acceptable")
	}

	nameCalled := "John"
	switch { //switch tanpa ekspresi, menggunakan case dengan kondisi boolean. lebih disarankan menggunakan if-else untuk kasus seperti ini, tapi switch juga bisa digunakan
	case nameCalled == "John":
		fmt.Println("Hello, John!")
	case nameCalled == "Doe":
		fmt.Println("Hello, Doe!")
	default:
		fmt.Println("Go Ahead!, john is not here")
	}
}
