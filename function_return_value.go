package main

import "fmt"

func greeting(name string) string {
	return "Morning sir, " + name
}

func main() {
	fmt.Println(greeting("Richards"))
}
