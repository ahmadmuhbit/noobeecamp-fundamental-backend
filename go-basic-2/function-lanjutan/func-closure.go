package main

import "fmt"

func main() {
	calculate := func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Println(calculate(1, 3))
}
