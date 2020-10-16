package main

import "fmt"

func Sekolah() (string, string)  {
	return "SMK 1 Majalengka", "Majalengka"
}

func main() {
	var name, address = Sekolah()
	fmt.Println(name)
	fmt.Println(address)
}
