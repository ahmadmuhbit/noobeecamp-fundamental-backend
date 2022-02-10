package main

import "fmt"

func main() {
	fruits := []string{"Grape", "Banana", "Apple"}

	for index, fruit := range fruits {
		fmt.Printf("Index %d adalah buah %s\n", index, fruit)
	}
}
