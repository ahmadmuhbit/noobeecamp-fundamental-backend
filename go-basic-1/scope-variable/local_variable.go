package main

import "fmt"

func main() {

	num := 2

	if num > 1 {
		text := "Lebih besari dari 1"
	} else {
		text := "Lebih kecil atau sama dengan 1"
	}

	fmt.Println(text) // Error : undefined :text

}
