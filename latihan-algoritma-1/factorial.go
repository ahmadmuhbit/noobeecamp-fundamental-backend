package main

import (
	"fmt"
)

func factorial_recursion(x int) (y int) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println(factorial_recursion(5))
}
