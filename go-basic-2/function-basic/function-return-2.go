package main

import "fmt"

func main() {
	fmt.Println(combineString("Hello", "World"))
}

func combineString(str1 string, str2 string) string {
	return str1 + " " + str2
}
