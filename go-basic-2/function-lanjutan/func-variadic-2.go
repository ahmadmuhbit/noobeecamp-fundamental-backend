package main

import "fmt"

func main() {
	var nums = []int{1, 8, 29, 3, 4, 5, 2, 9, 7}
	total := sum(nums...)
	fmt.Println(total)
}

func sum(num ...int) (total int) {
	fmt.Println(num)

	for _, n := range num {
		total += n
	}
	return total
}
