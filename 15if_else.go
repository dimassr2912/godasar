package main

import "fmt"

func main() {
	name := "Rama"

	if name == "Rama" { // Kondisi (Harus bertipe data boolean)
		fmt.Println("Hello Rama")
	} else if name == "Dimas" {
		fmt.Println("Hello Dimas")
	} else {
		fmt.Println("Hello boleh kenalan?")
	}

	// Short statement / variable temporary
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Silahkan")
	}
}

/**
 * If else: Kata kunci digunakan untuk percabangan
 *	Percabangan: Eksekusi program tertentu ketika kondisi terpenuhi
 * If mendukung short statement
 */
