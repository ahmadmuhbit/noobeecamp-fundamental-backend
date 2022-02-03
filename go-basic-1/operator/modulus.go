package main

import "fmt"

func main() {
	num := 10

	modulus4 := num % 4
	fmt.Printf("Modulus 4 : %d\n", modulus4)

	modulus3 := num % 3
	fmt.Printf("Modulus 3 : %d\n", modulus3)

	modulus5 := num % 5
	fmt.Printf("Modulus 5 : %d\n", modulus5)
}
