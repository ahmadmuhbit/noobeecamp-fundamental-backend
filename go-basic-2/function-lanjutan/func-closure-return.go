package main

import "fmt"

func main() {
	nilai := []int{5, 2, 8, 6, 10, 8, 9, 5, 8}
	kkm := 6

	total, nilaiLulus := graduate(nilai, kkm)
	fmt.Println("Total yang lulus :", total)
	fmt.Println("Nilai yang lulus :", nilaiLulus())
}

func graduate(nilaiArr []int, limit int) (int, func() []int) {
	var studentGraduate []int

	for _, arr := range nilaiArr {
		if arr >= limit {
			studentGraduate = append(studentGraduate, arr)
		}
	}

	return len(studentGraduate), func() []int {
		return studentGraduate
	}
}
