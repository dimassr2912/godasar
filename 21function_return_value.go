package main

import "fmt"

func getHello(name string) string {
	return "Hello" + name
}

func main() {
	result := getHello("Rama")
	fmt.Println(result)
}

/**
 * Return: Keyword agar function bisa mengembalikan data
 */
