package code2

/*
Soal
====================
Dalam sebuah kantong, terdapat beberapa bola.
Ada 3 warna bola, yaitu Merah, Kuning, dan Biru.

Tampilkanlah jumlah bola paling banyak

Diketahu :
Bola = Merah, Biru, Biru, Kuning, Merah
Warna = Merah, Biru, Kuning

Ditanya :
Bola apa yang paling banyak?
Dan berapa jumlahnya?

Jawab :
Warna merah = 2
Warna biru = 2
Warna kuning = 1

Karena merah dan biru sama, maka outputnya adalah :
merah : 2, biru : 2
*/
func GetBallsByColor(balls []string, colors []string) map[string]int {
	// lengkapi code berikut
	total := make(map[string]int)
	count := make(map[string]int)
	/*
		{
			'merah' : 2,
			'biru' : 2
		{
	*/

	max := 0

	for _, color := range colors {
		c := 0
		for _, ball := range balls {
			if color == ball {
				c++
			}
		}
		if max < c {
			max = c
		}

		total[color] = c
	}

	for color, tot := range total {
		if tot == max {
			count[color] = tot
		}
	}

	return count

}
