package main

import "fmt"

func main() {
	cetak()
}

func cetak() {
	fmt.Println("Fungsi cetak() dipanggil")
	hello()
}

func hello() {
	fmt.Println("Hello World")
}
