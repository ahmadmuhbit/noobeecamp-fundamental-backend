package main

import "fmt"

func main() {

	time := "am"

	switch time {
	case "pm":
		fmt.Println("Malam")
	case "am":
		fmt.Println("Pagi")
	default:
		fmt.Println("Maaf, waktunya salah!")
	}

}
