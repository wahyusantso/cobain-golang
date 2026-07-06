package main

import "fmt"

func main() {
	person := map[string]string{ //definisikan map dengan tipe data string untuk key dan value
		"name":    "John Doe",
		"address": "123 Main St",
		"city":    "Anytown",
		"age":     "30",
	}

	fmt.Println(person["name"])              //akses nilai dengan key "name"
	person["email"] = "john.doe@example.com" //tambahkan key "email" dengan nilai "john.doe@example.com"
	fmt.Println(person["email"])             //akses nilai dengan key "email"
	fmt.Println(person)                      //cetak seluruh map

	book := make(map[string]string) //buat map kosong dengan fungsi make
	book["title"] = "The Great Gatsby"
	book["author"] = "F. Scott Fitzgerald"
	fmt.Println(book) //cetak seluruh map book

	delete(book, "author") //hapus key "author" dari map book
	fmt.Println(book)      //cetak seluruh map book setelah penghapusan
}
