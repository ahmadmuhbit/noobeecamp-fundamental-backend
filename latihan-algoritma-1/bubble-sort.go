package main

import "fmt"

func main() {
	// 0(n^2)
	// 3 2 1
	// 2 1 3
	// 1 2 3

	a := []int{3, 2, 1}
	BubbleSort(a)
	fmt.Println(a)
}

func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
