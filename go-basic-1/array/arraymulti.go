package main

import "fmt"

func main() {

	var fruits = [2][3]string{
		[3]string{"Apple", "Mango", "Banana"},
		[3]string{"Orange", "Grape", "Avocado"},
	}

	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])

}
