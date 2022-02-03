package main

import "fmt"

func main() {

	num := 2

	if num > 2 {
		// kondisi ini akan dipanggil jika
		// syarat terpenuhi
		fmt.Println("Num lebih besar dari dua")
	} else {
		// kondisi ini akan dipanggil jika
		// kondisi pertama tidak terpenuhi
		fmt.Println("Num lebih kecil dan sama dari 2")
	}

	// akan dijalankan setelah program didalam selesai
	fmt.Println("Program selesai.")
}
