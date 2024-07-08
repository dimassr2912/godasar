package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	result := sumAll(1, 2, 3, 4, 5)
	fmt.Println(result)

	// Jika ada kasus slice yang harus dimasukkan ke paramater varags
	numbers := []int{10, 10, 10, 10, 10}
	result2 := sumAll(numbers...)
	fmt.Println(result2)
}

/**
 * Parameter yang berada di akhir, memiliki kemampuan dijadikan varargs
 * 	Varargs bisa menerima lebih dari 1 input
 * 	Varargs dianggap sebagai slice
 * Keuntungannya tidak perlu mengirim argument slice saat memanggil fungsi
 * Jika memiliki variable slice, diberi ...
 */
