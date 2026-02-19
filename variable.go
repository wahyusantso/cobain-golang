package main

import "fmt"

func main() {
	var name string //deklarasi variable dengan tipe data string dan tipe var
	name = "Wahyu Santoso"
	fmt.Println(name)

	name = "Budi"
	fmt.Println(name)

	address := "Jakarta" //deklarasi variable dengan tipe :=, go akan otomatis mendeteksi tipe data dan bersifat var
	fmt.Println(address)

	var ( //deklarasi multiple variable
		firstName = "Wahyu"
		lastName  = "Santoso"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
