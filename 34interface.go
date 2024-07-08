package main

import "fmt"

// Deklarasi interface
type HasName interface {
	GetName() string // Method
}

func sayHello3(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// Struct yang mengimplementasikan dari interface HasName
type Person struct {
	Name string
}

// Implementasi method GetName untuk tipe Person (Harus mengikuti kontrak interface HasName)
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Rama"}
	sayHello3(person)

	animal := Animal{Name: "Anjing"}
	sayHello3(animal)
}

/**
 * Interface: Tipe data abstract, tidak memiliki implementasi langsung
 *	Kumpulan definisi method yang tidak memiliki isi (definisi saja)
 *	Biasa digunakan untuk kontrak
 *		Kontrak diimplementasi dalam struct
 *	Merupakan tipe data yang value zero adalah nil
 */
