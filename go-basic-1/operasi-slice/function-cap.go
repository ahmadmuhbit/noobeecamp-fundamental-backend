package main

import "fmt"

func main() {
	animals := []string{"Cat", "Dog", "Cow"}
	anim1 := animals[0:2]
	anim2 := animals[1:3]

	fmt.Println("====[ANIMALS]====")
	fmt.Println(animals)
	fmt.Println("Capacity :", cap(animals))
	fmt.Println("Len \t :", len(animals))
	fmt.Println("")

	fmt.Println("====[ANIM1]====")
	fmt.Println(anim1)
	fmt.Println("Capacity :", cap(anim1))
	fmt.Println("Len \t :", len(anim1))
	fmt.Println("")

	fmt.Println("====[ANIM2]====")
	fmt.Println(anim2)
	fmt.Println("Capacity :", cap(anim2))
	fmt.Println("Len \t :", len(anim2))
	fmt.Println("")

}
