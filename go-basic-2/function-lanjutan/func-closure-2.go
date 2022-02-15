package main

import "fmt"

func main() {
	sorting := func(arr []int) []int {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if arr[i] < arr[j] {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
		return arr
	}
	var arr = []int{3, 9, 5, 1, 10, 4, 2}
	fmt.Println("Awal : ", arr)
	fmt.Println("Setelah diurut :", sorting((arr)))
}
