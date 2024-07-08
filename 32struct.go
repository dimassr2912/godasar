package main

import "fmt"

// Deklarasi struct
type Customer struct {
	// Terdiri dari 3 field / property / attribute
	Name, Address string
	Age           int
}

func main() {
	// Deklarasi instance dari struct Customer
	var s1 Customer

	// Mengisi field
	s1.Name = "Dimas"
	s1.Address = "Kalasan"
	s1.Age = 25

	fmt.Println(s1)
	fmt.Println(s1.Name)
	fmt.Println(s1.Address)
	fmt.Println(s1.Age)

	// Struct literal
	s2 := Customer{
		Name:    "Dimas",
		Address: "Kalasan",
		Age:     25,
	}

	fmt.Println(s2)

	// Harus sesuai urutan jika menggunakan ini
	s3 := Customer{"Rama", "Kalasan", 25}
	fmt.Println(s3)

}

/**
 * Struct: Template data yang digunakan untuk menggabungkan nol atau lebih tipe data
 *	Struct kumpulan dari field / property / attribute (variable)
 * Struct tidak bisa langsung digunakan, tetapi harus membuat data / object dari struct (object struct)
 * Jika property tidak diisi akan diberi nilai default tiap tipe data
 */
