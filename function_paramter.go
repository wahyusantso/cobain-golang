package main

import "fmt"

func sayHello(name, from string) { // deklare param function dengan tipe yang sama
	fmt.Println("Hello", name, ", your from", from)
}

func main() {
	sayHello("John do", "Poland")
}
