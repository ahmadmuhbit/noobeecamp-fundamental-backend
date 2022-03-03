package main

import "fmt"

func main() {
	var test interface{}
	test = "Hello bro"
	fmt.Println(test)

	test = map[string]string{
		"name": "NBID",
	}
	fmt.Println(test)

	test = 3
	fmt.Println(test)

	test = false
	fmt.Println(test)
}
