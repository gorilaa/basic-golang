package main

import (
	"fmt"
	"math"
)

func main() {
	// Round selalu mencari pembulatan ke yang terdekat
	fmt.Println(math.Round(3.3)) // pembulatan ke bawah
	fmt.Println(math.Round(4.7)) // pembulatan ke atas

	// Floor selalu melakukan pembulatan ke bawah
	fmt.Println(math.Floor(3.3)) // pembulatan ke bawah
	fmt.Println(math.Floor(4.7)) // pembulatan ke bawah

	// Ceil selalu melakukan pembulatan ke atas
	fmt.Println(math.Ceil(3.3)) // pembulatan ke atas
	fmt.Println(math.Ceil(4.7)) // pembulatan ke atas

	// Min and Max
	fmt.Println(math.Min(10, 20)) // nilai terendah
	fmt.Println(math.Max(10, 20)) // nilai terbesar
}
