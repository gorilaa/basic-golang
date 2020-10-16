package main

import "fmt"

func factorialLoop(value int) int  {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecrusive(value-1)
	}
}

func main() {
	// factorial Loop
	factorialLoop := factorialLoop(5)
	fmt.Println(factorialLoop)

	// factorial Recrusive
	factorialRecrusive := factorialRecrusive(5)
	fmt.Println(factorialRecrusive)

	fmt.Println(5 * 4 * 3 * 2 * 1)
}
