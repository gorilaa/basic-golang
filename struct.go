package main

import "fmt"

type Personal struct {
	Name, Address string
	Age 		  int
	statusMerried bool
}

func main() {
	// case ine
	var dataPersonalOne Personal
	dataPersonalOne.Name = "Adam Lesmana"
	dataPersonalOne.Address = "Majalengka"
	dataPersonalOne.Age = 27
	dataPersonalOne.statusMerried = true

	fmt.Println(dataPersonalOne)

	//case two
	dataPersonalTwo := Personal{
		Name : "Ganda Saputra",
		Address: "Bandung",
		Age: 29,
		statusMerried: false,
	}

	fmt.Println(dataPersonalTwo)
}
