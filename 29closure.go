package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}

/**
* Closure: Kemampuan function berinteraksi dengan data-data disekitarnya
* 	Akan pusing ketika kode kompleks, tiba-tiba counter berubah nilainya
 */
