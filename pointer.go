package main

import "fmt"

type Address struct {
	City, Province, Country string
}
func main() {
	addressOne := Address{"Majalengka", "Jawa Barat", "Indonesia"}
	//addressTwo := addressOne
	addressTwo := &addressOne

	//change city addressTwo
	addressTwo.City = "Sumedang"
	/**
	addressOne tidak berubh saat addressTwo merubah City
	jika ingin merubah apa yang di referensikan gunakan & , contoh addressTwo := &addressOne
	 */
	fmt.Println(addressOne)

	fmt.Println(addressTwo)
}
