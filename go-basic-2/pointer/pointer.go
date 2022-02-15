package main

import "fmt"

func main() {
	nama := "NooBee"
	namaPointer := &nama

	fmt.Println(nama)
	fmt.Println(namaPointer)
	fmt.Println(*namaPointer)
}
