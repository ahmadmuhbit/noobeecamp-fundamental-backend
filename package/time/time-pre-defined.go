package main

import (
	"fmt"
	"time"
)

func main() {
	myTimeStr := "2022-01-16T15:04:05Z"

	myTime, err := time.Parse(time.RFC3339, myTimeStr)
	if err != nil {
		fmt.Println("err :", err)
	} else {
		fmt.Println("Oke :", myTime)
	}
}
