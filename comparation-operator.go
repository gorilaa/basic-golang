package main

import "fmt"

func main() {
	var nameOne  = "Adam"
	var nameTwo  = "Lesmana"
	var nameTree = "Adam"

	var resultOne bool = nameOne == nameTree
	var resultTwo bool = nameOne == nameTwo

	fmt.Println(resultOne)
	fmt.Print(resultTwo)
}
