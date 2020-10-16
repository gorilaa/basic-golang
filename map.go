package main

import "fmt"

func main() {
	person := map[string]string{
		"name" : "Adam Lesmana Ganda Saputra",
		"address" : "Majalengka - Jawa Barat",
	}

	person["title"] = "Software Engineer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	delete(person, "title")

	fmt.Println(person)
}
