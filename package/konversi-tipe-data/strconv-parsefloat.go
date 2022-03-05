package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12.34"

	num, err := strconv.ParseFloat(numStr, 32)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}
}
