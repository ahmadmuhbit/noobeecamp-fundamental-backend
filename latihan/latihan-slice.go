package main

import "fmt"

func main() {

	// array
	animals := [...]string{"Cat", "Dog", "Pinguin", "Chicken", "Snake"}

	mammals := animals[0:2]
	notMammals := animals[2:5]
	haveLegs := animals[:4]

	fmt.Println(mammals)
	fmt.Println(notMammals)
	fmt.Println(haveLegs)

	println("")

	mammals[1] = "Cow"
	notMammals[0] = "Bird"

	fmt.Println(mammals)
	fmt.Println(notMammals)
	fmt.Println(haveLegs)

}
