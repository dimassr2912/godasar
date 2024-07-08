package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

/**
 * Main function adalah fungsi yang akan dijalankan ketika program dijalankan
 * Main function diletakkan di dalam main package
 * Setiap file harus memiliki package
 * Setiap module harus memiliki 1 package main
 * Import digunakan untuk import package lain ke dalam file
 * 	Package fmt (format) adalah Standard library golang
 * 		Function Println() digunakan untuk menampilkan output tulisan
 * Pada saat compile, akan dicari file yang memiliki main function sesuai dengan OS yang digunakan
 */
