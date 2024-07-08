package main

import "fmt"

func endApp2() {
	fmt.Println("END APP")
	message := recover() // Letakkan disini karena defer akan selalu dipanggil saat error
	fmt.Println("Terjadi error:", message)
}

func runApp2(error bool) {
	defer endApp2()
	if error {
		panic("ERROR")
	}
	// message := recover() Tidak diletakkan disini karena saat panic error, kode selanjutnya tidak dieksekusi
}

func main() {
	runApp2(true)
}

/**
 * Recover: function yang digunakan untuk menangkap data panic
 *	Program akan tetap berjalan
 * Jika memanggil recover() akan mengambil isi dari panic()
 */
