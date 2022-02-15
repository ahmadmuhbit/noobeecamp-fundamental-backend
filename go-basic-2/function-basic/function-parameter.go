package main

import "fmt"

func main() {
	cetak("Hello World!")
	//cetak(3) // akan error
}

func cetak(text string) {
	fmt.Println(text)
}
