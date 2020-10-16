package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])m")

	fmt.Println(regex.MatchString("adm"))
	fmt.Println(regex.MatchString("aim"))
	fmt.Println(regex.MatchString("aDm"))
}
