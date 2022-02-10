package main

import "fmt"

func main() {

	person := map[string]string{
		"name": "NooBee",
		"job":  "Programmer",
	}

	for key, value := range person {
		fmt.Printf("key : %s, value : %s \n", key, value)
	}

}
