package main

import "fmt"

func main() {
	var names = []string{"Rey", "Jo", "NooBee"}
	var ages = []int{10, 13, 20}

	for _, name := range names {
		cetakNama(name)
	}

	for _, age := range ages {
		cetakUmur(age)
	}
}

func cetakNama(text string) {
	fmt.Println(text)
}

func cetakUmur(age int) {
	fmt.Println(age, "tahun")
}
