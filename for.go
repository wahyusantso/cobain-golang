package main

import "fmt"

func main() {
	counter := 2

	for counter <= 10 {
		// fmt.Println("Counter:", counter)
		counter += 2
	}

	for decCounter := 10; decCounter >= 2; decCounter -= 2 { //for dengan short statement untuk menghitung mundur dari 10 ke 0
		// fmt.Println("DecCounter:", decCounter)
	}

	names := []string{"John", "Doe", "Jane", "Smith"}
	for _, name := range names { //for range untuk iterasi slice names, index untuk mendapatkan indeks dan name untuk mendapatkan nilai
		// fmt.Printf("Index: %d, Name: %s\n", index, name)
		fmt.Printf("Name: %s\n", name)
	}
}
