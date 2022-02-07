package main

import "fmt"

func main() {

	animals := [4]string{"Cat", "Dog", "Chicken", "Bird"}

	legs4 := animals[0:2]
	legs2 := animals[2:4]

	fmt.Println(legs4)
	fmt.Println(legs2)

}
