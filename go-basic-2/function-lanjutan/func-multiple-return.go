package main

import "fmt"

func main() {

	var sisi float32 = 4

	luas, keliling := calculate(sisi)
	fmt.Println("Luas & keliling dari persegi dengan panjang sisi", sisi)
	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)

}

func calculate(sisi float32) (float32, float32) {
	luas := sisi * sisi
	keliling := sisi * 4
	return luas, keliling
}
