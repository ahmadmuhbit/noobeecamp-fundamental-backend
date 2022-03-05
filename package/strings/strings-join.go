package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Nama saya NooBeeID"

	stringSplit := strings.Split(string1, " ")
	stringJoin := strings.Join(stringSplit, "-")

	fmt.Println(string1)
	fmt.Println(stringSplit)
	fmt.Println(stringJoin)
}
