package main

import "fmt"

func logging() {
	fmt.Println("Selesai")
}

func runApplication() {
	defer logging() // Akan dipanggil diakhir
	fmt.Println("Run application")
}

func main() {
	runApplication()
}

/**
 * Defer: function yang dijadwalkan dieksekusi setelah sebuah function selesai dieksekusi
 *	Dieksekusi walaupun terjadi error
 * Jika function error, maka akan langsung berhenti eksekusi functionnya
 */
