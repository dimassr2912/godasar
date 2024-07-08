package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Hello ", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Hello", i)
	}

	// Perulangan tanpa for range
	names := []string{
		"Dimas", "Setyawan", "Ramadhansyah",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	for i, _ := range names {
		fmt.Println(names[i])
	}
}

/**
 * For: Keyword yang digunakan untuk perulangan
 *	Selama kondisi bernilai true, blok kode akan selalu dieksekusi
 * 	for <init statement>; <kondisi>; <post statement>
 *		init statement hanya dieksekusi sekali
 *		post statement akan dieksekusi setelah blok program dijalankan
 * For bisa digunakan untuk iterasi data collection (array, slice, map)
 */
