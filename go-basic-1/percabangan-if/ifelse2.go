package main

import "fmt"

func main() {

	age := 18
	var hasKTP bool

	if age > 17 {
		hasKTP = true
	} else {
		hasKTP = false
	}

	fmt.Printf("Apakah saya mempunyai KTP? : %v", hasKTP)

}
