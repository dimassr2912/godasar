package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Dimas" // Tanpa perlu diinisialisasi
	lastName = "Ramadhansyah"
	return // Bisa tanpa disebutkan variablenya
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}

/**
 * Bisa juga tidak diisii sehingga menggunakan default value
 */
