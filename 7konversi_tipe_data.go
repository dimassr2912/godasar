package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	// Number Overflows
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name string = "Dimas"
	var d = name[0]
	var dString = string(d)

	fmt.Println(name)
	fmt.Println(d)
	fmt.Println(dString)
}

/**
 * Number overflows: Jika tipe data tidak bisa menampung nilai / melebihi kapasitas, maka akan dihitung ke belakang, maka akan menjadi -32768
 */
