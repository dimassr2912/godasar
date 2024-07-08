package main

import "fmt"

func Ups() interface{} {
	return 1
	// return "Ups"
	// return true
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}

/**
 * Interface kosong: Interface yang tidak memliki deklarasi method
 *	Membuat seluruh tipe data menjadi implementasinya
 * 	Memiliki alias any
 * 	Contoh: fmt.Println(), panic(), recover()
 */
