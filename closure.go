package main

import "fmt"

func main() {
	name := "Adam Lesmana"
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		/**
		- Bisa merubah nilai variable yang di deklasrasikan di atas nya
		- Nilai counter akan bertambah jika function digunakan
		- Variable yang di buat di dalam lingkup function tidak bisa di akses oleh yang lain
		 */
		name = "Ganda Saputra"
		counter++
	}

	increment()
	increment()
	fmt.Println(name)
	fmt.Println(counter)
}
