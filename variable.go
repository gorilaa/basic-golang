package main

import "fmt"

func main() {
	//singe
	var name , address string
	var (
		boyName string
		bornAddress string
	)

	name = "Adam Lesamana"
	address = "Majalengka"
	fmt.Println(
		"Nama Saya",
		name,
		"Alamat Saya",
		address,
	)

	name = "Adam Saputra"
	address = "Bandung"
	fmt.Println(
		"Nama Saya",
		name,
		"Alamat Saya",
		address,
	)

	boyName = "Gavin Khawarzmi"
	bornAddress = "Majalengka"
	fmt.Println("My name ",boyName, bornAddress)

	//short variable
	namaLengkap := "Adam Lesmana Ganda Saputra"
	fmt.Println(namaLengkap)
}