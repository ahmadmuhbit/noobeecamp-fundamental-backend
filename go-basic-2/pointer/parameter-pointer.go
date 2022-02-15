package main

import "fmt"

func main() {
	myName := "NooBee"
	fmt.Println("Nama awal :", myName)

	changeName(&myName, "NooBeeID")
	fmt.Println("Nama setelah diubah :", myName)

}

func changeName(name *string, newName string) {
	fmt.Println("Change name from", *name, "to", newName)
	*name = newName
}
