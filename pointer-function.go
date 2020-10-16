package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddressCountry(address *Address)  {
	address.Country = "Indonesia"
}

func main() {
	addressOne := Address{"Majalengka", "Jawa Barat", ""}
	changeAddressCountry(&addressOne) // gunakan pointer "&"

	fmt.Println(addressOne)
}
