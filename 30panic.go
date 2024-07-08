package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}
func main() {
	runApp(false)
	runApp(true)
}

/**
 * Panic: function yang digunakan untuk menghentikan program
 *
 */
