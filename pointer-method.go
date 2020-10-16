package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Merried() {
	man.Name = "Mr. "+ man.Name
}
func main() {
	name := Man{Name: "Adam Lesmana"}
	name.Merried()

	// tidak akan berubah jika data nya as variable , harus dirubah jadi pointer
	// gunakan bintang(*) pada function Merried()
	fmt.Println(name)
}
