package main

import "fmt"

type Filter func(string) string

func spamHelloWithFilter(name string, filter Filter)  {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

func spamFilter(name string) string {
	if name == "Biawak" {
		return "..."
	}
	return name
}

func main() {
	spamHelloWithFilter("Adam", spamFilter)
	spamHelloWithFilter("Biawak", spamFilter)
}
