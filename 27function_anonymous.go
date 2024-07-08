package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	// Anonymous function
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Rama", blacklist)
	registerUser("Rama", func(name string) bool {
		return name == "Anjing"
	})
}

/**
 * Anonymous function: Function tanpa nama yang dibuat langsung di variable / parameter
 */
