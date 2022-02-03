package main

import "fmt"

func main() {

	name := new(string)
	*name = "NooBeeID"

	fmt.Printf("Ini adalah pointer : %s dan ini adalah valuenya : %s", name, *name)

}
