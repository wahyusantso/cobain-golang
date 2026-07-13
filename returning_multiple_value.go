package main

import "fmt"

func getAddress(address string, no int) (string, int) { //return sesuai tipe data return
	return address, no
}

func main() {
	address, no := getAddress("Jl. Semeru No.", 24)
	fmt.Println(address, no)
}
