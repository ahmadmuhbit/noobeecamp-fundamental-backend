package main

import "fmt"

func main() {
	var car = make(map[string]string)
	car["name"] = "BMW"
	car["color"] = "Black"

	message := combineString(car["name"], car["color"])
	cetak(message)

}

func combineString(str1 string, str2 string) string {
	return "Mobil " + str1 + " Berwarna " + str2
}

func cetak(str3 string) {
	fmt.Println(str3)
}
