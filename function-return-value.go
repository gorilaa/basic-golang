package main

import "fmt"

func Belajar(namaPelajaran string) string  {
	fmt.Println("Ini adalah mata pelajaran:", namaPelajaran)
	return namaPelajaran
}

func main() {
	var namaPelajaran = Belajar("Bahasa Indonesia")
	fmt.Println(namaPelajaran)
}
