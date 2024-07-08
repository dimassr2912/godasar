package main

import (
	"fmt"
)

func getFullName() (string, string) {
	firstNama := "Dimas"
	lastName := "Rama"
	return firstNama, lastName
}

func main() {
	first, last := getFullName()
	fmt.Println(first, last)

	// Menghiraukan return value
	_, lastName := getFullName()
	fmt.Println(lastName)
}

/**
 * Function bisa return banyak value
 * Jika ingin menghiraukan return value, gunankan _
 * Biasanya digunakan pada tipe: map, slice, struct
 */
