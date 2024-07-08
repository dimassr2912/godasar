package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()

	// Bisa terjadi panic

	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// Gunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

/**
 * Type assertions: Ubah tipe data menjadi yang diinginkan
 *	Memastikan bahwa tipe data benar
 */
