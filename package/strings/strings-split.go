package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Nama saya NooBeeID"

	stringSplit := strings.Split(string1, " ")

	fmt.Println(string1)
	fmt.Println(stringSplit)
}