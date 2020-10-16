package main

import "fmt"

func random() interface{} {
	return false
}

func main() {
	var result interface{} = random()
	
	//switch 
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown")

	}
}