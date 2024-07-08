package main

import (
	"errors"
	"fmt"
)

func pembagian(pembilang, penyebut int) (int, error) {
	if penyebut == 0 {
		return 0, errors.New("Penyebut 0")
	} else {
		return pembilang / penyebut, nil // Nil karena interface bisa diisi nil
	}
}

func main() {
	hasil, err := pembagian(10, 2)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}

/**
 * Golang memiliki interface error: type error interface{Error() string}
 *	Ada package error, jadi tidak perlu membuat struct
 * errors.New(): Untuk membuat pesan error
 */
