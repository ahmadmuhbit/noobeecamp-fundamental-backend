package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 = 1
	var num2 = "1"

	/* jika mencoba kode dibawah, maka akan muncul error
	   invalid operation: num1 == num2 (mismatched types int and string)
	   hal ini terjadi karena num1 berupa int dan num2 berupa string.
	**/
	//fmt.Println("Hasil perbandingan :", num1 == num2) // ini akan error

	num1Str := strconv.Itoa(num1)
	fmt.Println("Hasil perbandingan :", num1Str == num2)
}
