package structt

import "fmt"

// struct
type Customer struct{
	Name,Address string
	Age int
}

func Struct() Customer {
	newCustomer := Customer{
		Name : "rofi",
		Address: "Permata",
		Age : 23,
	}

	return newCustomer
}

// membuat variabel dari struct
func VariabelStruct() string{
	// cara 1
	var rofi Customer
	rofi.Name = "Rofi"
	rofi.Address = "Permata"
	rofi.Age = 30

	return rofi.Name
}

// Struct literal, beberapa cara untk membuat struct
func StructLiterals() Customer{
	// cara 2
	joko := Customer{
		Name : "Joko",
		Address : "Indonesia",
		Age : 30,
	}

	// cara 3 ==> kurang recomend pa eko, karena bisa jadi error karena takut salah pemasukan data, jika data nya banyak
	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	return joko
}

// Struct
/*
	1. Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
	2. Struct biasanya representasikan data dalam program aplikasi yang kita buat
	3. Data di struct di simpan dalam field
	4. Sederhananya struct adalah kumpulan field
*/

// Membuat Data Struct ==> membuat variabel nya
/*
	1. Struct adalah template data atau prototype data
	2. Struct tidak bisa lansung digunakan 
	3. Namun kita bisa membuat data/object dari struct yang telalh kita buat
*/

// Struct Literals
/*
	1. sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
*/