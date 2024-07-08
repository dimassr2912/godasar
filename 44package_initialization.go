package main

import (
	"fmt"
	"go-dasar/database"
	_ "go-dasar/internal" // Menggunakan blank indetifier karena ingin memanggil init saja tanpa function lain
)

func main() {
	fmt.Println(database.GetDatabase())
}

/**
 * init: Membuat function yang langsung diakses ketika package digunakan
 *	Cocok digunakan untuk berkomunikasi dengan database
 * Blank indentifier: Jika ingin menjalankan init function di package tanpa memanggil apapun (Tambahkan _)
 */
