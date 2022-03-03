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
	// jika salah satu ada yang tidak diisi, kita tau ini akan menghasilkan error.
	// Nah untuk menghandle itu, kita akan membuat error kita sendiri
	// inilah yang dimaksud dengan error yang expected
	if s.Name == "" {
		return errors.New("Nama tidak boleh kosong")
	}

	if len(s.Name) <= 3 {
		return errors.New("Nama harus lebih dari 3")
	}

	if s.Height == 0 {
		return errors.New("Tinggi tidak boleh kosong")
	}

	return nil
}

func main() {
	student := Student{Name: "NooBee"}

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
