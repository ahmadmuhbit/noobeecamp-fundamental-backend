package main

import "fmt"

var name string = "NooBee"

func main() {

	fmt.Println(name, "dipanggi dari fungsi main")
	sayHello()

}

func sayHello() {
	fmt.Println(name, "dipanggil dari fungsi sayHello")
}
