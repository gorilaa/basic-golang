package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Adam Lesmana", "Adam")) // cek kata
	fmt.Println(strings.Split("Adam Lesmana Ganda Saputra", " ")) // memotong menjadi slice
	fmt.Println(strings.ToLower("Adam Lesmana Ganda Saputra")) // huruf kecil semua
	fmt.Println(strings.ToUpper("Adam Lesmana Ganda Saputra")) // Huruf besar semua
	fmt.Println(strings.Trim("     Adam Lesmana Ganda Saputra    ", " ")) // membersihkan spasi
	fmt.Println(strings.ReplaceAll("Adam Lesmana Ganda Saputra", "Lesmana", "ganteng"))
}
