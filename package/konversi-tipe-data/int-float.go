package main

import "fmt"

func main() {
	num1 := 10
	float1 := float32(num1)

	fmt.Printf("num1 :%d\n", num1)
	fmt.Printf("float1 :%f\n", float1)
	fmt.Printf("Casting float to int :%d\n", int(float1))
}
