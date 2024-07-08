package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// Pass by value
	address1 := Address{
		"Jogja", "DIY", "Indonesia",
	}

	address2 := address1 // copy value
	address2.City = "Solo"

	fmt.Println(address1) // Tidak ikut berubah
	fmt.Println(address2) // Berubah menjadi solo

	// Pass by reference
	address3 := &address1
	address3.City = "Surabaya"
	fmt.Println(address3)
	fmt.Println(address1)

	// 	*address3 = Address{"Surabaya", "Jatim", "Indonesia"}

	// 	fmt.Println(address1)
	// 	fmt.Println(address3)

	// 	// Function new()
	// 	alamat1 := new(Address)
	// 	alamat2 := alamat1

	// 	alamat1.City = "Jakarta"
	// 	fmt.Println(alamat1)
	// 	fmt.Println(alamat2)

	// 	// variable pointer
	// 	// var number *int

	// 	var numberA int = 4
	// 	var numberB *int = &numberA

	// 	fmt.Println(numberA)
	// 	fmt.Println(&numberA)

	// 	fmt.Println(numberB)
	// 	fmt.Println(*numberB)

	// 	// Akan ikut berubah karena memiliki alamat memori yang sama
	// 	numberA = 5
	// 	fmt.Println(numberA)
	// 	fmt.Println(&numberA)

	// 	fmt.Println(numberB)
	// 	fmt.Println(*numberB)

	// 	// Parameter pointer
	// 	number := 4
	// 	change(&number, 10)
	// 	fmt.Println(number)
	// }

	// func change(original *int, value int) {
	// 	*original = value
}

/**
 * Variable di golang passing by value bukan passing by reference
 * 	Jika mengirim variable ke function / method / variable lain, yang dikirim duplikasi value
 * Pointer: kemampuan membuat reference lokasi data memory yang sama tanpa duplikasi (pass by reference)
 * Untuk membuat variable pointer ke variable lain, gunakan &
 *	Variable pointer berarti: variable yang berisi alamat memori suatu nilai, ditandai dengan *
 * Referencing: Mengambil nilai pointer dari variable biasa dengan menambah &
 * Deferencing: Mengambil nilai asli dari variable pointer dengan menambah *
 * Keyword new: Membuat variable pointer dengan tipe data tertentu
 */
