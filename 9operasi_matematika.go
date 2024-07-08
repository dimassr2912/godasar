package main

import "fmt"

func main() {
	// Operasi matematika
	var a int = 10
	var b = 10
	c := a + b
	fmt.Println(c)

	// Augmented assignment
	a += 10 // a = a + 10
	fmt.Println(a)

	// Unary operator
	fmt.Println(-a)

}

/**
 * Digunakan untuk operasi tipe data number
 * Operator matematika:
 *	+, -, *, /, %
 * Augmented assignment"
 *	+=, -=. *=, /=, %=
 * Unary operator:
 *	++, --, -, +, !
 */
