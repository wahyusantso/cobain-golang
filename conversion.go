package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64) // akan minus, karena int16 tidak bisa menampung nilai 32768
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Wahyu Santoso"
	var e = name[1]
	var eString = string(e) // akan mengubah byte menjadi string

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
