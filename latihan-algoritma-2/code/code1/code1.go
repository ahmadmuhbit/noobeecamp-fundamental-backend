package code1

/*
Soal
===============
Pada sebuah kue, terdapat beberapa lilin
dengan tinggi yang berbeda.

Saat kamu meniup lilin tersebut, maka yang kamu tiup adalah
lilin paling tinggi. Jadi yang akan padam adalah lilin paling tinggi.

Diketahui :
Tinggi lilin = 3,2,4,4,1

Ditanya :
Berapa lilin yang akan padam ?

Jawab :
Pada lilin tersebut, yang paling tinggi adalah 4
Dan jumlah lilin yang mempunyai tinggi 4 adalah 2 buah

Maka, lilin yang akan padam sejumlah 2 buah
*/
func GetCandles(candles []int) int {
	// lengkapi code berikut ...
	count := 0

	if len(candles) == 0 {
		return 0
	}

	max := candles[0]
	// proses untuk nilai maksimum
	for _, candle := range candles {
		if max < candle {
			max = candle
		}
	}

	// proses untuk melihat, ada berapa lilin dengan
	// tinggi maksimum
	for _, candle := range candles {
		if candle == max {
			count++
		}
	}

	return count
}
