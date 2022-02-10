package main

import "fmt"

func main() {

	nilai := map[string]int{
		"Agama":      80,
		"Matematika": 90,
		"Olahraga":   70,
		"Design":     80,
	}

	fmt.Println("===== Nilai Awal =====")
	fmt.Println(len(nilai))
	fmt.Println(nilai)

	fmt.Println("\n\n===== Nilai Akhir =====")
	delete(nilai, "Design")
	fmt.Println(len(nilai))
	fmt.Println(nilai)

}
