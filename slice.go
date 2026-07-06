package main

import "fmt"

func main() {
	colors := [...]string{"Blue", "Yellow", "Black", "Red", "Green", "Purple"}
	//  [pointer, Kapasitas]
	slice1 := colors[2:6]
	slice2 := colors[0:]
	slice3 := colors[:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
