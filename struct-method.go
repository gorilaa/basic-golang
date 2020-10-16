package main

import "fmt"

type Customer struct {
	Name string
	Age int
	statusMerried bool
}

func (customer Customer) sayLove() {
	fmt.Println("I love u", customer.Name)
}
func main() {
	adam := Customer{Name: "Adam"}
	adam.sayLove()
}
