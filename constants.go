package main

import "fmt"

func main() {
	//single conts
	const name string = "Adam Lesmana Ganda"
	const address 	  = "Bandung"

	//group const
	const (
		boyName string = "Gavin Kaharizmi"
		addressName = "Majalengka"
	)

	fmt.Println(name)
	fmt.Println(address)

	fmt.Println("Nama Saya", boyName, addressName)
}
