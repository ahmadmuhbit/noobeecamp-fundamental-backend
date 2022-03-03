package main

import "fmt"

func main() {
	defer cetak("Text 1")
	defer cetak("Text 2")
	cetak("Text 3")
	defer cetak("Text 4")
	cetak("Text 5")
}

func cetak(text string) {
	fmt.Println(text)
}
