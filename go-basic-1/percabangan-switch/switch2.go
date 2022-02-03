package main

import (
	"fmt"
	"time"
)

func main() {

	// cara normal

	/*switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Weekend")
	case time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekdays")
	}*/

	// cara singkat

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekdays")
	}

}
