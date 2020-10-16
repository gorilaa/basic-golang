package main

import (
	"fmt"
	"golang-dasar/helper"
	"golang-dasar/env"
)

func main() {
	//import dari package helper
	result := helper.SayHelloBosku("Adam")

	fmt.Println(result)

	//import dari package env
	env := environment.Application;
	fmt.Println(env)
}
