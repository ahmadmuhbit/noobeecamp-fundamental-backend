package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("Waktu sekarang:", now.String())
	fmt.Println("Tahun sekarang:", now.Year())
	fmt.Println("Bulan sekarang:", now.Month())
	fmt.Println("Tanggal sekarang:", now.Day())
}

// Gimana caranya nampilin bulan dalam bentuk int : misal februari = 2
