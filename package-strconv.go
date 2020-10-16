package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println("Hasilnya adalah ", boolean)
	} else {
		fmt.Println("Terjadi Error:", err.Error())
	}

	// decimal = 10
	// binary = 2
	// oktal = 8
	// hexadecimal 16
	// 64 adalah tipe jenis int , misal int8 = 8, int16 = 16 dst
	number , err := strconv.ParseInt("10000", 10, 64)
	if err == nil {
		fmt.Println("Hasilnya adalah ", number)
	} else {
		fmt.Println("Terjadi Error:", err.Error())
	}

	value := strconv.FormatInt(1000000, 2)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("100000") // parsing string to integer
	fmt.Println(valueInt)

	valueString := strconv.Itoa(100000)
	fmt.Println(valueString)
}
