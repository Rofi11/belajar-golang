package structt

import "fmt"

// struct function/method
func (function Customer) SayHello(name string) {
	fmt.Println("Hello", name, "My Name is", function.Name)
}

// struct
func StructMethod(function Customer, name string) string{
	function.SayHello(name)
	return "test"
}

// Struct Mehod
/*
	1. Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk fucntion
	2. Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan -akan sebuah struct memiliki function
	3. Method adalah function
*/