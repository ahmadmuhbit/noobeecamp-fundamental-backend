package main

import "fmt"

type Calculate interface {
	Divide() float32
	Multiply() float32
}

// sebuah struct yang akan mengimplement method dari interface
type Num1 struct {
	Num     float32
	Divider float32
}

// proses mendefinisikan struct tadi kedalam interface
func NewNum1(num float32, divider float32) Calculate {
	return &Num1{Num: num, Divider: divider}
}

// implement method pada interface
func (n Num1) Divide() float32 {
	return n.Num / n.Divider
}

// implement method pada interface
func (n Num1) Multiply() float32 {
	return n.Num * n.Num
}

func main() {
	num := NewNum1(4.0, 2.0)
	fmt.Println(num.Divide())
}
