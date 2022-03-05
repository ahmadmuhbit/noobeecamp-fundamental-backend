package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	myTimeStr := "2021-01-16 10:00:00"
	
	myTime, err := time.Parse(layout, myTimeStr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(myTime.String())
	}
}
