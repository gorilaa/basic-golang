package main

import "fmt"

func main() {
	name := "Adam"

	if name == "Adam" {
		fmt.Println("Bener bro namanya", name)
	} else if name == "Lesmana" {
		fmt.Println("Ini Baru", name)
	} else {
		fmt.Println("Nama lu sape ?")
	}

	// Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Ini Bener")
	} else {
		fmt.Println("Ini Salah")
	}
}
