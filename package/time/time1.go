package main

import (
	"fmt"
	"time"
)

func main() {
	var timeDateLocal time.Time = time.Date(2022, 1, 16, 10, 15, 0, 0, time.Local)
	var timeDateUTC time.Time = time.Date(2022, 1, 16, 10, 15, 0, 0, time.UTC)
	var timeNow time.Time = time.Now()

	fmt.Printf("Time Date Local : %v\n", timeDateLocal)
	fmt.Printf("Time Date UTC : %v\n", timeDateUTC)
	fmt.Printf("Time Now : %v\n", timeNow)
}
