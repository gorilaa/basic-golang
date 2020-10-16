package main

import "fmt"

func startApp(error bool)  {
	defer endApp()
	if error {
		panic("Aplikasi error ")
	}

	fmt.Println("Aplikasi dimulai")
}

func endApp()  {
	// Recover itu selalu dilakukan di akhir
	message := recover();
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

	fmt.Println("Aplikasi selesai")
}

func main() {
	startApp(true)
}