package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name" : "Adam Lesmana Ganda Saputra",
			"address" : "Bandung",
		}
	}
}
func main() {
	name := newMap("baron")
	fmt.Println(name)
}
