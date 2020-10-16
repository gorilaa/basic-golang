package main

import "fmt"

func main() {
	for i :=0; i <50; i++ {
		if i == 25 {
			break
		}

		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke:", i)
	}
}
