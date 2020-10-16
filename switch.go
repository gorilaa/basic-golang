package main

import "fmt"

func main() {
	name := "Adam"

	switch name {
	case "Adam":
		fmt.Println("Ini ", name)
	case "Lesmana":
		fmt.Println("Ini si ", name)
	default:
		fmt.Println("Lu Siape ?")
	}

	//short statement
	switch length := len(name); length > 3 {
	case true:
		fmt.Println("Sesuai")
	case false:
		fmt.Println("Gagal")
	}
}
