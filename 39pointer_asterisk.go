package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		"Jogja", "DIY", "Indonesia",
	}
	address3 := &address1
	fmt.Println(address1)
	fmt.Println(address3)

	// address3 = &Address{"Kalasan", "Jogja", "Indonesia"} // Berbeda karena pindah pointer
	// fmt.Println(address1)
	// fmt.Println(address3)

	*address3 = Address{"Solo", "Jateng", "Indonesia"}
	fmt.Println(address1) // Ikut berubah
	fmt.Println(address3)
}

/**
 * Variable biasa bisa diambil nilai pointernya dengan menambah tanda &
 * Nilai asli variable pointer bisa diambil, dengan menambahkan tanda *
 */
