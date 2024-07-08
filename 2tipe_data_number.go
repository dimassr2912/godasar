package main

import "fmt"

func main() {
	fmt.Println("Satu ", 1)
	fmt.Println("Dua", 2)
	fmt.Println("Tiga koma lima", 3.5)
}

/**
 * Tipe data golang terbagi menjadi 4 kategori:
 * Basic type: number, boolean, string
 * Aggregate type: array, struct
 * Reference type: slice, pointer, map, function, channel
 * Interface type: Interface
 * Tipe data number:
 * 	Integer (Non-desimal)
 *		uint / unsign integer (Bilangan cacah)
 *			uint8, uint16, uint32, uint64
 *		int (Bilangan bulat)
 *			int8, int16, int32, int64
 * 	Floating point (Desimal)
 *		float32, float64
 * Tentukan tipe data sesuai nilai untuk menghindari boros memori karena semakin besar tipe data semakin besar penggunaan memori
 * Alias
 *	byte = uint8
 * 	rune = int32
 * 	int = int32 / int64 (Sesuai nilai)
 *	uint = uint32 / uint64 (Sesuai nilai)
 * Zero value: 0 (non-desimal), 0.0 (desimal)
 */
