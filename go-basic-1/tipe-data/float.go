package main

import "fmt"

func main() {

	var contohDecimal float32 = 1.66
	fmt.Printf("%f\n", contohDecimal)   // default, 6 angka dibelakang koma
	fmt.Printf("%.1f\n", contohDecimal) // dibulatkan
	fmt.Printf("%.3f\n", contohDecimal) // tiga angka dibelakang koma

}
