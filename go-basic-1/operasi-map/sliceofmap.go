package main

import "fmt"

func main() {

	students := []map[string]string{
		map[string]string{"name": "NooBee", "major": "Teknik Komputer"},
		map[string]string{"name": "Jovie", "major": "Sistem Informasi"},
		map[string]string{"name": "Reyhan", "major": "Teknik Informatika"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "berada di jurusan", student["major"])
	}

}
