package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Method
func (customer Customer) sayHello(name string) { // sayHello adalah method yang terkait dengan tipe Customer
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var rama Customer

	rama.Name = "Dimas"
	rama.Address = "Kalasan"
	rama.Age = 25

	rama.sayHello("Rama") // Harus dipanggil dengan object customer karena method bukan lagi function
}

/**
 * Struct bisa digunakan seperti tipe data lainnya yaitu sebagai paramater function
 * Method adalah function yang menempel pada struct
 * Untuk akses method gunakan dari object
 *	(customer Customer) adalah receiver dari method
 */
