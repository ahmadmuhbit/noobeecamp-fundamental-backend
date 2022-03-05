package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "1234"

	num, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}
}
