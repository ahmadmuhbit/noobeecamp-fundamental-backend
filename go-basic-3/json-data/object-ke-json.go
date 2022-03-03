package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Class string
}

func main() {
	users := []User{
		User{Name: "NooB", Class: "A"},
		User{Name: "Java", Class: "B"},
		User{Name: "Docker", Class: "C"},
	}

	fmt.Println("Bentuk dasar :", users)

	usersJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bentuk JSON :", string(usersJSON))
}
