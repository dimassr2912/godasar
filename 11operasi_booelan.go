package main

import "fmt"

func main() {
	var nilaiUjian = 90
	var nilaiAbsen = 80

	var lulusUjian = nilaiUjian > 80
	var lulusAbsen = nilaiAbsen > 70

	var lulus bool = lulusAbsen && lulusUjian
	fmt.Println(lulus)
}

/**
 * Operator boolean:
 * &&, ||, !
 */
