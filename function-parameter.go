package main

import "fmt"

func Register(name string) {
	fmt.Println("Nama saya adalah:", name)
}

func main() {
	name := "Adam Lesmana"
	Register(name)
}
