package main

import "fmt"

func main() {
	num1 := 10
	num2 := 5

	total := sum(num1, num2)
	cetak(total)

	// bisa disingkat dengan
	cetak(sum(30, 50))
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func cetak(arg interface{}) {
	fmt.Println(arg)
}
