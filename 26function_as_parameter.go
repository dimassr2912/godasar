package main

import "fmt"

// Membuat alias / type declarations
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

// Harus memiliki contract parameter dengan tipe data string dan return tipe data string
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Rama", spamFilter)
}

/**
 * Function bisa digunakan sebagai paramter function lain
 */
