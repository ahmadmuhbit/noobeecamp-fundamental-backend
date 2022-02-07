package main

func main() {

	var fruits [3]string // jumlah data yang ditampung ada 3

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange" // batas data, yaitu panjang array -1

	fruits[3] = "Grape" // akan error, karena melebihi kapasitas array

}
