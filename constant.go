package main

import "fmt"

func main() {
	const name = "Wahyu Santoso"
	fmt.Println(name)

	// name = "Budi" //error karena constant tidak bisa diubah

	const ( //multiple deklarasi constant
		firstName = "Wahyu"
		lastName  = "Santoso"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
