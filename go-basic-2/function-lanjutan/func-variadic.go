package main

import "fmt"

func main() {
	fmt.Println(sum(1, 8, 29, 3, 4, 5, 2, 9, 7))
}

func sum(num ...int) (total int) {
	fmt.Println(num)

	for _, n := range num {
		total += n
	}
	return total
}
