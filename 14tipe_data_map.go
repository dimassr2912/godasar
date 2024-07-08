package main

import "fmt"

func main() {
	// Deklarasi
	var person map[string]string = map[string]string{}

	// Assign
	person["Name"] = "Rama"
	person["Address"] = "Kalasan"
	fmt.Println(person)

	// Inisialisasi
	var person2 = map[string]string{
		"Name":    "Dimas",
		"Address": "Kalasan",
	}

	fmt.Println(person2)
	fmt.Println(person2["Name"])
	fmt.Println(person2["Address"])
	// Jika tidak ada, maka akan diisi default sesuai tipe data map
	fmt.Println(person2["Salah"])

	// Make
	book := make(map[string]string)
	book["Author"] = "Dimas"
	book["Title"] = "Golang"

	fmt.Println(book)

	delete(book, "Title")
	fmt.Println(book)

	// Cek item dengan key tertentu
	var value, isExist = book["Author"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item not exist")
	}

	// Kombinasi slice & map
	//[]map[string]int artinya slice yang tiap elemennya adalah map
	var animals = []map[string]int{
		{"Anjing": 1, "Ekor": 1},
		{"Kucing": 3}, // Tiap elemen dapat memiliki jumlah key yang berbeda
	}

	for _, animal := range animals {
		fmt.Println(animal["Anjing"])
	}
}

/**
 * Map: Berisi kumpulan data yang sama yang indexnya bisa selain number
 * 	Tipe data kumpulan key-value (pair)
 *		key tidak boleh sama dengan yang lain
 *	Jumlah data map boleh sebanyak-banyaknya
 * Function di map:
 *	len(map): Mendapatkan jumlah data di map
 * 	map[key]: Ambil data berdasarkan key
 *	map[key] = value: Ubah data map dengan key
 *	make(map[TypeKey]TypeValue): Membuat variable map
 * 	delete(map, key): Hapus data di map dengan key
 */
