package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari", "February", "Maret", "April", "Mei", "Juni",
		"July", "Agustus", "September", "Oktober", "November", "Desember",
	}

	fmt.Println(months)

	var sliceOne = months[4:7]
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))
	fmt.Println(sliceOne)


	sliceTwo := append(sliceOne, "Nambah nih bro")

	fmt.Println(sliceTwo)

	months[5] = "Bulan apa ini"
	fmt.Println(months)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
