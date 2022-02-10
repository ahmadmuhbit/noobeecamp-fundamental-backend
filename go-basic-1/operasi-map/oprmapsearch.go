package main

import "fmt"

func main() {

	nilai := map[string]int{
		"Agama":      80,
		"Matematika": 90,
		"Olahraga":   70,
		"Design":     80,
	}

	key := "Agama"
	val, isExist := nilai[key]

	if isExist {
		fmt.Println(key, "is exist with value", val)
	} else {
		fmt.Println(key, "is not exist")
	}

}
