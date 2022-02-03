package main

import "fmt"

func main() {

	var age uint8 = 256     // maksimal uint8 adalah 255
	fmt.Printf("%d\n", age) // Error : constant 256 overflows uint8

}
