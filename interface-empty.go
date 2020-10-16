package main

import "fmt"

func upsTwo() interface{} {
	return "Kosong"
}
func upsOne(i int) interface{} {
	switch i {
	case 1:
		return 1
	case 2:
		return true
	default:
		return "Ups ga ada"
	}
}
func main()  {
	upsTwo := upsTwo()

	fmt.Println(upsTwo)

	upsOne := upsOne(3)

	fmt.Println(upsOne)

}