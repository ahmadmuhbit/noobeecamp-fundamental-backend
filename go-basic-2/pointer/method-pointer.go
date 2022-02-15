package main

import "fmt"

type Car struct {
	Name  string
	Color string
}

func (c *Car) setName(newName string) {
	fmt.Println("Try to change car from ", c.Name, " name to =>", newName)
	c.Name = newName
}

func (c Car) changeName(newName string) {
	fmt.Println("Try to change car from ", c.Name, " name to =>", newName)
}

func main() {
	car := Car{Name: "BMW", Color: "White"}
	fmt.Println(car)

	car.setName("Civic")
	fmt.Println(car)

	car.changeName("Chevrolet")
	fmt.Println(car)
}
