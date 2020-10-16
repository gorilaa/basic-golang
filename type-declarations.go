package main

import "fmt"

func main() {
	type noKtp string
	type merried bool

	var nomorKtpAdam noKtp = "3211000043457876"
	var status merried = false
	fmt.Println(nomorKtpAdam)
	fmt.Println(status)
}
