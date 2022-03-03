package main

import (
	"fmt"
	"os"
)

func main() {
	names := []string{"NooB", "Bee", "Go", "NodeJS"}

	for _, name := range names {
		if name == "Go" {
			os.Exit(1)
		}
		cetak(name)
	}
}

func cetak(text string) {
	fmt.Println(text)
}
