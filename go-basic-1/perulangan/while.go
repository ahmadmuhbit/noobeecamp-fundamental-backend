package main

import "fmt"

func main() {
	var counter int = 0 // initialization

	for counter < 5 { // condition
		fmt.Println("Counter :", counter)

		counter++
	}
	fmt.Println("Counter berhenti")
}
