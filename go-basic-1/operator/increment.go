package main

import "fmt"

func main() {
	i := 0
	fmt.Println(i++) // ini akan error, karena i++ merupakan statements

	// post increment
	i++ // ini tidak error, karena merupakan sebuah statement
	fmt.Println(i)

	// pre increment
	++i // ini akan error, karena pada go hanya mengizinkan notasi secara postfix
}