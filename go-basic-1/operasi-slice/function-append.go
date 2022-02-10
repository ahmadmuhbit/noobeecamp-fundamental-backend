package main

import "fmt"

func main() {
	animals := []string{"Cat", "Dog", "Cow"}
	animSmall := animals[0:2]

	anim1 := append(animals, "Sheep")
	anim2 := append(animSmall, "Ant")

	fmt.Println("====[ANIMALS]====")
	fmt.Println("Capacity :", cap(animals))
	fmt.Println("Len \t :", len(animals))
	fmt.Println("Data \t :", animals)
	fmt.Println("")

	fmt.Println("====[ANIMALS SMALL]====")
	fmt.Println("Capacity :", cap(animSmall))
	fmt.Println("Len \t :", len(animSmall))
	fmt.Println("Data \t :", animSmall)
	fmt.Println("")

	fmt.Println("====[ANIM1]====")
	fmt.Println("Capacity :", cap(anim1))
	fmt.Println("Len \t :", len(anim1))
	fmt.Println("Data \t :", anim1)
	fmt.Println("")

	fmt.Println("====[ANIM2]====")
	fmt.Println("Capacity :", cap(anim2))
	fmt.Println("Len \t :", len(anim2))
	fmt.Println("Data \t :", anim2)
	fmt.Println("")

}
