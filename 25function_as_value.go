package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Dimas"))
}

/**
 * Function adalah first class citizen
 *	Function bisa disimpan dalam variable sebagai value
 *	Function dianggap sebagai tipe data
 */
