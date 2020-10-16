package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName  = "Bambang"
	middleName = "Joko"
	lastName   = "Pamungkas"

	return
}

func main() {
	firstName, middleName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
}
