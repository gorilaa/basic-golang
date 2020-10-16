package main

import "fmt"

func goodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	results := goodBye
	fmt.Println(results("Adam"))
}