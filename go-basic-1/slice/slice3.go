package main

import "fmt"

func main() {

	animals := [4]string{"Cat", "Dog", "Chicken", "Bird"}

	anim1 := animals[0:2] // semua element dimulai dari index 0 sampai index sebelum 2
	anim2 := animals[1:4] // semua element dimulai dari index 1 sampai index sebelum 4
	anim3 := animals[:]   // semua element
	anim4 := animals[1:]  // semua element dimulai dari index 1 (sampai akhir element)
	anim5 := animals[:3]  // semua element dimulai dari index 0 sampai index sebelum 3

	fmt.Println(anim1)
	fmt.Println(anim2)
	fmt.Println(anim3)
	fmt.Println(anim4)
	fmt.Println(anim5)

}
