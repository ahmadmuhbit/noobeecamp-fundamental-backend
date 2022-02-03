package main

import "fmt"

func main() {

	nilai := 80
	var grade string

	if nilai > 80 {
		grade = "A"
	} else if nilai > 70 {
		grade = "B"
	} else if nilai > 60 {
		grade = "C"
	} else if nilai > 50 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Println("Grade saya adalah : " + grade)

}
