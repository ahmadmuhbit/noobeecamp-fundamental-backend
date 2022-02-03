package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Subuh - Pagi")
	case t.Hour() > 12:
		fmt.Println("Siang - Malam")
	default:
		fmt.Println("Tidak tau jam berapa")
	}

}
