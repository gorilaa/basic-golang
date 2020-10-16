package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	// dibalik bos nilai nya
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Adam", 22},
		{"Lesmana", 20},
		{"Ganda ", 28},
		{"Saputra", 24},
	}

	sort.Sort(UserSlice(users)) // dikonversi menjadi alias

	fmt.Println(users)

}
