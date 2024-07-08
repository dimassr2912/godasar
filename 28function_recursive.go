package main

import "fmt"

func recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursive(value-1)
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	// result := recursive(5)
	// fmt.Println(result)
	result := factorialLoop(5)
	fmt.Println(result)

	result2 := recursive(5)
	fmt.Println(result2)
}

/**
 * Recursive function: Function yang memanggil dirinya sendiri
 * Contoh kasus: Factorial
 */
