package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10 - 5
	var c = a * b

	fmt.Println(c)

	//augment assigned
	var i = 10
	i += 10
	fmt.Println(i)

	//unary
	i++
	fmt.Println(i)
}