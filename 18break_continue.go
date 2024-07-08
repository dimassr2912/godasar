package main

import "fmt"

func main() {
	// Break
	for i := 0; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println("Hello", i)
	}

	// Continue
	for j := 0; j < 10; j++ {
		if j%2 == 1 {
			continue
		}
		fmt.Println(j)
	}
}

/**
 * Break dan continue: Kata kunci yang bisa digunakan untuk perulangan
 *	Break: Untuk menghentikan seluruh perulangan
 * 	Continue: Untuk menghentikan perulangan yang berjalan dan melanjutkan ke perulangan selanjutnya
 */
