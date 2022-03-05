package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Nama saya NooBeeID"

	fmt.Println("Text asli :", string1)
	fmt.Println("Ubah 1 huruf a :", strings.Replace(string1, "a", "o", 1))
	fmt.Println("Ubah 2 huruf a :", strings.Replace(string1, "a", "o", 2))
	fmt.Println("Ubah 3 huruf a :", strings.Replace(string1, "a", "o", 3))
	fmt.Println("Ubah semua huruf a :", strings.Replace(string1, "a", "o", -1))
}
