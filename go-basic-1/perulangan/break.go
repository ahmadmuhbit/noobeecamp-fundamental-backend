package main

import "fmt"

func main() {
	var isTen bool
	counter := 0

	for {
		if counter == 10 {
			isTen = true
			break
		}
		isTen = false
		counter++
	}

	fmt.Println("Counter :", counter, "Is Ten :", isTen)

}
