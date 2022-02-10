package main

import "fmt"

type Parent struct {
	Nama string
	Umur int
}

type Student struct {
	OrangTua Parent
	Nama     string
	Umur     int
	Kelas    string
}

func main() {

	parent1 := Parent{
		Nama: "Andi",
		Umur: 50,
	}

	student1 := Student{
		OrangTua: parent1,
		Nama:     "Budi",
		Umur:     11,
		Kelas:    "6A",
	}

	student2 := Student{
		OrangTua: Parent{
			Nama: "Jojo",
			Umur: 51,
		},
		Nama:  "Bilqis",
		Umur:  10,
		Kelas: "5B",
	}

	fmt.Println(student1)
	fmt.Println(student2)

	fmt.Printf("Orang tua %s Bernama %s\n", student1.Nama, student1.OrangTua.Nama)
	fmt.Printf("Orang tua %s Bernama %s\n", student2.Nama, student2.OrangTua.Nama)

}
