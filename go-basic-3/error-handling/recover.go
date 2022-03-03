package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name   string
	Height int
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("Nama tidak boleh kosong")
	}

	if len(s.Name) <= 3 {
		panic("Nama harus lebih besar dari 3")
	}

	if s.Height == 0 {
		return errors.New("Tinggi tidak boleh kosong")
	}

	return nil
}

func catchError() {
	// panic tadi akan masuk kedalam recover()
	err := recover()
	if err != nil {
		cetak(fmt.Sprintf("Muncul error nih %v", err))
	}
}

func main() {
	defer catchError()
	student := Student{Name: "NB"}

	err := student.Validate()
	if err != nil {
		cetak(err.Error())
	} else {
		cetak("Hello, " + student.Name)
	}
}

func cetak(text string) {
	fmt.Println(text)
}
