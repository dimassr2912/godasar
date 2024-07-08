package main

import "fmt"

func main() {
	var number = 4
	fmt.Println("before changeByValue:", number) // 4

	changeByValue(number)
	fmt.Println("after changeByValue :", number) // 4 (tidak berubah)
}

func changeByValue(original int) {
	original = 10
}
