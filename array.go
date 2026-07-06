package main

import "fmt"

func main() {
	names := [4]string{"Inaya", "Husen", "Edward", "Finia"}
	flexible := [...]int{01, 0, 24, 5} //tidak terbatas jumlahnya
	var color [2]string
	color[0] = "red"
	color[1] = "blue"

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(len(flexible))
	fmt.Println(color)
	color[1] = "Locked"
	fmt.Println(color)
}
