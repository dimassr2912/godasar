package main

import "fmt"

func main() {
	// Pointer: 4 | length: 2, capacity: 2
	// slice bisa di assign dengan 3 cara: append(), inisialisasi, make()
	var names = [...]string{"Dimas", "Setyawan", "Ramadhansyah", "Ira", "Meidiana", "Putri"}

	// Slice dibentuk dari array dengan teknik 2 index | Nilai yang didapat adalah reference
	var slice1 []string = names[4:6]
	var slice2 = names[:3]
	slice3 := names[3:]
	slice4 := names[:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	// 1 index | Nilai yang didapat adalah copy
	slice5 := names[2]
	fmt.Println(slice5)

	// Function slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "Sabtu baru"
	daysSlice1[1] = "Minggu baru"

	// array days akan berubah karena daysSlice1 reference ke array days
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur")
	daysSlice2[0] = "Ups"

	// Tidak mengubah array days karena kapasitas penuh, sehingga membuat array baru
	fmt.Println(days)
	fmt.Println(daysSlice2)

	// Make
	newSlice := make([]int, 2, 5)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, 3)
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	// copy
	fromSlice := days[:]
	toSlice := make([]string, len(days), cap(days))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}

/**
 * Tipe Data Slice: Potongan dari data array dan reference elemen array
 * 	Ukuran dari slice bisa berubah
 * 	Slice akses seluruh atau sebagian dari array
 *	Jika data array penuh, membuat array baru dan semua data dipindahkan ke array yang baru
 *	Memiliki 3 data yang penting: pointer, length, capacity
 * 		Pointer: Penunjuk data pertama
 *		Length: Panjang dari slice
 * 		Capacity: Kapasitas dari slice (capacity >= length)
 *	Slice bisa dihasilkan dari manipulasi array atau slice lain
 * 		Slice bisa dibuat dari array dengan 2 index
 *			array[low:high]: Dimulai dari index low hingga high
 *			array[low:]: Dimulai dari index low hingga akhir array
 *			array[:high]: Dimulai dari index 0 hingga high
 *			array[:]: Seluruh index
 * Function slice:
 *	len(slice): Mendapatkan panjang slice
 *	cap(slice): Mendapatkan kapasitas slice
 *	append(slice, data): Menambahkan data ke posisi terakhir
 *		Jika kapasitas sudah penuh, maka membuat array baru
 * 	make([]typeData, length, capacity): Membuat variable slice baru
 *	copy(destination, source): Menyalin dari source ke destination
 */
