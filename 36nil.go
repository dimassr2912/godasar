package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data["name"])
	}
}

/**
 * nil: Sebuah nilai yang berarti nilai kosong
 * 	Bisa digunakan pada tipe data: interface, function, pointer, slice, map, channel
 */
