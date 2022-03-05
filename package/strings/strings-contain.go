package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Nama saya NooBeeID"
	cari := "NooBeeID"

	isExist := strings.Contains(string1, cari)
	fmt.Printf("Apakah string : %s ada didalam text %s ? %v\n", cari, string1, isExist)
}
