package main

import "fmt"

func main() {
	var name string                      // Deklarasi
	name = "Dimas Setyawan Ramadhansyah" // Assign value
	fmt.Println(name)

	// Re-assign
	name = "Rama"
	fmt.Println(name)

	// Inisialisasi
	var age = 25
	fmt.Println(age)

	// Tanpa var dan tipe data
	address := "Kalasan"
	fmt.Println(address)

	address = "Jogja"
	fmt.Println(address)

	// Multiple variable
	var (
		firstName = "Dimas Setyawan"
		lastName  = "Ramadhansyah"
	)
	fmt.Println(firstName, lastName)
}

/**
 * Variable: tempat menyimpan data
 *	Digunakan untuk akses data yang sama dimanapun
 * Jika inisialisasi langsung, bisa tanpa tipe data, tetapi tipe data ditentukan oleh nilainya
 * Var tidak wajib jika saat inisialisasi, tetapi ditambah menggunakan :=
 * Jika variable tidak digunakan, akan error
 */
