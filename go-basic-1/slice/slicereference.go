package main

import "fmt"

func main() {

	animals := [4]string{"Cat", "Dog", "Chicken", "Bird"}

	anim1 := animals[0:3]
	anim2 := animals[2:4]

	fmt.Println(animals) // [Cat, Dog, Chicken, Bird]
	fmt.Println(anim1)   // [Cat, Dog, Chicken]
	fmt.Println(anim2)   // [Chicken, Bird]

	anim1[1] = "Cow"
	fmt.Println(animals) // [Cat, Cow, Chicken, Bird]
	fmt.Println(anim1)   // [Cat, Cow, Chicken]
	fmt.Println(anim2)   // [Chicken, Birds]

	anim2[0] = "Pinguin"
	fmt.Println(animals) // [Cat, Cow, Pinguin, Bird]
	fmt.Println(anim1)   // [Cat, Cow, Penguin]
	fmt.Println(anim2)   // [Pinguin, Bird]

}
