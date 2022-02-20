package main

import (
	"fmt"
)

func main() {
	fmt.Println(insertionSort([]int{5, 2, 4, 6, 1, 3})) // Rata-rata: O(k*n)
	fmt.Println(insertionSort([]int{1, 2, 3, 4, 5, 6})) // Terbaik: O(n)
	fmt.Println(insertionSort([]int{6, 5, 4, 3, 2, 1})) // Terburuk: O(n*n)
}

// insertionSort mengurutkan array/slice yang diberikan menggunakan algoritma INSERTION SORT
// 1. Saya mengulangi dari kiri ke kanan
// 2. Sisi kiri selalu diurutkan
// 3. Lalu mengambil setiap elemen dan memasukkannya ke posisi yang benar
func insertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i > -1 && A[i] > key {
			A[i+1] = A[i]
			i -= 1
		}
		A[i+1] = key
	}
	return A
}
