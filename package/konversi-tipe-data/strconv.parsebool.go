package main

import (
	"fmt"
	"strconv"
)

func main() {
	bString := "true"

	b, err := strconv.ParseBool(bString)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(b == true)
	}
}
