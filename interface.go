package main

import "fmt"

type Kids struct {
	Name string
	Age int
}
type PersonalKids interface {
	GetName() string
}

func helloKid(kids PersonalKids)  {
	fmt.Println("Hello kids", kids.GetName())
}

func (kids Kids) GetName() string {
	return kids.Name
}

func main() {
	kidData := Kids{Name: "Gavin Khawarizmi", Age: 27}

	helloKid(kidData)
}
