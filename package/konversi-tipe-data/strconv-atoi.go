package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 = 1
	var num2 = "1"

	num2Int, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num1 == num2Int)
	}
}
