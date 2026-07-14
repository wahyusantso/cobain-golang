package main

import "fmt"

func sumAll(numbers ...int) int { //like aray yang menyimpan banyak data, namun hanya bisa di parameter akhir
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(20, 10, 20)) //bebas masukin data sepuas nya selama integer

	//contoh case dimana ada slice yang dikirim ke varargs
	numbers := []int{10, 20, 10}
	fmt.Println(sumAll(numbers...)) //disini titik tiga digunakan untuk membongkar slice menjadi satu per satu, jika tidak akan error.
}
