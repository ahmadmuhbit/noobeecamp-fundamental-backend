package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20

	// kita ingin mengubah num1, dengan cara menjumlahkan num1 dan num2

	// cara normal
	// num1 = num1 + num2
	// fmt.Printf("Cara normal : %d\n", num1)

	// cara singkat
	num1 += num2
	fmt.Printf("Cara singkat : %d\n", num1)
}
