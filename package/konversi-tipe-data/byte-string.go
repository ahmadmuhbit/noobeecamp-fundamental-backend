package main

import "fmt"

func main() {
	str1 := "Haii"
	byte1 := []byte(str1)

	fmt.Printf("str1 : %v\n", str1)
	fmt.Printf("byte1 : %v\n", byte1)
	fmt.Printf("Case dari byte1 ke string : %v\n", string(byte1))
	fmt.Printf("Case dari byte ke string : %v\n", string(65))
}
