package main

import "fmt"

func identity() (name, address string, age int) { //definisi variabel di return function
	name = "Ricard Stefan"
	address = "Pekanbaru"
	age = 25

	return name, address, age
}

func main() {
	first, middle, last := identity()
	fmt.Println(first, middle, last)
}
