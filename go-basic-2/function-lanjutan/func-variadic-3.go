package main

import "fmt"

const (
	Plus  string = "+"
	Minus string = "-"
)

func main() {
	nums := []float32{1, 4, 2, 9}
	result := calculate(Plus, nums...)
	fmt.Println(result)
}

func calculate(operator string, numbers ...float32) (result float32) {
	for _, num := range numbers {
		if operator == Plus {
			result += num
		} else if operator == Minus {
			result -= num
		}
	}
	return result
}
