package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := true
	
	bString := strconv.FormatBool(b)

	fmt.Println(bString == "true")
}
