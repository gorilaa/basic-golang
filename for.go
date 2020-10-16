package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Ini counter ke :", counter)
		counter++
	}

	// for with statements
	for counter :=1; counter <= 10; counter++ {
		fmt.Println("Ini counter statement ke:", counter)
	}

	name := []string{"Adam", "Lesmana", "Ganda", "Saputra"}

	for i :=0; i < len(name); i++ {
		fmt.Println("Nama ke:", i,  name[i])
	}

	// for range statement

	for i, v := range name {
		fmt.Println("Ini adalah index ke:", i, "=", v);
	}
}
