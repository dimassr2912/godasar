package main

import "fmt"

func main() {
	type noKtp string

	var ktpRama noKtp = "11111"

	var contoh string = "22222"
	var contohKtp noKtp = noKtp(contoh)

	fmt.Println(ktpRama)
	fmt.Println(contohKtp)
}

/**
 * Type declaration: Membuat ulang tipe data baru dari tipe data yang sudah ada
 *	Biasa digunakan untuk alias agar mudah dimengerti
 */
