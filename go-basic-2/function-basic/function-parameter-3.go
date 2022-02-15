package main

import "fmt"

func main() {
	var names = []string{"Rey", "Jo", "NooBee"}
	var ages = []int{10, 13, 20}

	for _, name := range names {
		cetak(name)
	}

	for _, age := range ages {
		cetak(age)
	}

}

func cetak(arg interface{}) {
	fmt.Println(arg)
}
