package main

import "fmt"

func main() {
	name := "Rama"
	switch name {
	case "Dimas", "Setyawan", "Ramadhansyah":
		fmt.Println("Hello Dimas")
	case "Rama":
		fmt.Println("Hello Rama")
		fallthrough
	default:
		{
			fmt.Println("Hello boleh kenalan?")
			fmt.Println("Minta no.nya?")
		}
	}

	// Short statment
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Silahkan")
	}

	// Switch tanpa kondisi
	switch {
	case name == "Rama":
		fmt.Println("Hello Rama")
	case name == "Dimas":
		fmt.Println("Hello Dimas")
	default:
		{
			fmt.Println("Hello boleh kenalan?")
			fmt.Println("Minta no.nya?")
		}
	}
}

/**
 * Switch: Salah satu keyword percabangan yang lebih sederhana
 *	Digunakan untuk pengecekan pada satu variable saja (==)
 */
