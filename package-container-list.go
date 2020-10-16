package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Adam")
	data.PushBack("Lesmana")
	data.PushBack("Ganda")
	data.PushBack("Saputra")

	fmt.Println(data.Front().Value) // ambil data paling depan
	fmt.Println(data.Back().Value) // ambil data paling belakang

	/**
	DARI DEPAN KE BELAKANG
	e ke 1 membuat variable dan mengambil data yang paling depan
	e ke 2 memastikan datanya tidak sama dengan nil
	e ke 3 melanjutkan ke data selanjutnya sampai e nya bernilai nil
	 */
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	/**
	DARI BELAKANG KE DEPAN
	e ke 1 membuat variable dan mengambil data yang paling Belakang
	e ke 2 memastikan datanya tidak sama dengan nil
	e ke 3 melanjutkan ke data sebelumnya sampai e nya bernilai nil
	*/
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
