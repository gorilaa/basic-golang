package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	absensi := 80

	lulusNilai := nilaiAkhir > 80 // trye
	lulusaAbsensi := absensi > 80 // false

	var results bool = lulusNilai && lulusaAbsensi

	fmt.Println(results) // false

}
