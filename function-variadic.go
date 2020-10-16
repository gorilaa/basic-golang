package main

import "fmt"

func sumAll(numbers ...int) int  {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	numbersOne := sumAll(1, 2, 3, 4, 5)
	fmt.Println(numbersOne)

	// variadic with slice
	numbersTwo := []int{6, 7, 8, 9, 10}
	fmt.Println(sumAll(numbersTwo...))
}
