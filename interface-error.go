package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian error, karena nilai pembagi adalah")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	results, err := Pembagian(10,0);
	if err == nil {
		fmt.Println("Resuts:", results)
	}

	fmt.Println("Error:", err)
}