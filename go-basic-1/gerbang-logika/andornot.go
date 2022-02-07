package main

import "fmt"

func main() {

	gender := "male"
	age := 17

	var canWork bool

	// penggunaan AND
	if gender == "male" && age >= 18 {
		canWork = true
	} else {
		canWork = false
	}

	fmt.Println("Nilai canWork pada notasi AND", canWork)

	// penggunaan OR
	if gender == "male" || age >= 18 {
		canWork = true
	} else {
		canWork = false
	}

	fmt.Println("Nilai canWork pada notasi OR", canWork)

	// penggunaan NOT
	if gender != "male" {
		canWork = true
	} else {
		canWork = false
	}

	fmt.Println("Nilai canWork pada notasi NOT", canWork)

}
