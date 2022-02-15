package main

import "fmt"

type Student struct {
	Name string
}

func (s *Student) SetMyName(newName string) {
	s.Name = newName
}

func (s Student) CallMyName(newName string) {
	fmt.Println("Hello, My Name is", newName)
}

func main() {
	student := Student{Name: "NooBee"}
	fmt.Println(student)

	student.CallMyName("Ahmad Muhbit")
}
