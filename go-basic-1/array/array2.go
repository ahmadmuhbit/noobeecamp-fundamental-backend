package main

import "fmt"

func main() {

	// vertical array
	fruits := [3]string{
		"Apple",
		"Banana",
		"Orange", // wajib menambahkan koma di blok akhir
	}

	fmt.Println(fruits)

	// Horizontal Array
	cars := [3]string{"Honda", "Toyota", "Nissan"}

	fmt.Println(cars)

	// Tanpa Panjang Data
	phones := [...]string{
		"Iphone",
		"Samsung",
		"Oppo",
	}

	fmt.Println(phones)

	// Menggunakan keyword make
	banks := make([]string, 3)
	banks[0] = "BCA"
	banks[1] = "Mandiri"
	banks[2] = "BTN"

	fmt.Println(banks)

}
