package main

import "fmt"

func main() {
	var names[3]string
	names[0] = "Adam"
	names[1] = "Lesmana"
	names[2] = "Ganda Saputra"

	fmt.Println(names)

	var value = [3]int{
		90,80,70,
	}

	fmt.Println(value)
	// length array
	fmt.Println(len(value))
	// change array value
	value[0] = 1000
	fmt.Println(value)
}
