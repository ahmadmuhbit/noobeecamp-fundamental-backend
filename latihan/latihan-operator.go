package main

import "fmt"

func main() {
	// Latihan
	/*
		Ubah kode dibawah ini menjadi lebih singkat
	*/

	// sekarang
	umurSaya := 20
	umurKakak := 30
	fmt.Printf("Umur saya sekarang adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak sekarang adalah : %d\n", umurKakak)

	// tahun depan

	// cara normal
	//umurSaya = umurSaya + 1
	//umurKakak = umurKakak + 1

	// cara singkat
	//umurSaya += 1
	//umurKakak += 1

	//fmt.Printf("Umur saya tahun depan adalah : %d\n", umurSaya)
	//fmt.Printf("Umur kakak tahun depan adalah : %d\n", umurKakak)

	// 5 tahun lagi

	// cara normal
	//umurSaya = umurSaya + 5
	//umurKakak = umurKakak + 5

	// cara singkat
	umurSaya += 5
	umurKakak += 5

	fmt.Printf("Umur saya 5 tahun lagi adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak 5 tahun lagi adalah : %d\n", umurKakak)

}
