package main

import "fmt"

func main() {
	// Deklarasi
	var name [3]string

	// Re-assign
	name[0] = "Dimas"
	name[1] = "Setyawan"
	name[2] = "Ramadhansyah"

	// Akses data / ambil data
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// Inisialisasi
	var score = [3]int{
		90,
		80,
	}

	score[2] = 100
	fmt.Println(score)

	// Function array
	fmt.Println(len(score))
	fmt.Println(score[2])
	score[1] = 100

	// Jumlah data ditentukan nilai yang di assign
	var address = [...]string{
		"Dogongan", "Tirtomartani", "Kalasan",
	}

	fmt.Println(address)
}

/**
 * Array: Tipe data kumpulan data dengan tipe data yang sama
 * 	Perlu menentukan jumlah data yang bisa ditampung
 *	Daya tampung tidak bisa ditambah setelah dibuat
 * 	Default nilai tiap elemen array tergantung tipe datanya
 * Function array:
 *	len(array): Mendapatkan panjang array
 * 	array[index]: Mendapatkan data di posisi index
 *	array[index] = value: Modifkiasi data di posisi index
 */
