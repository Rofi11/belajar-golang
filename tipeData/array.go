package tipeData

import "fmt"

func Array() string {
	// membuat array
	// cara 1, yg masukin ke variabel setelah declare tipe data dan jumlah data
	var names [3]string
	names[0] = "eko"
	names[1] = "kurniawan"
	names[2] = "khannedy"

	fmt.Println((names[0]))

	// cara 2, membuat array langsung saat declarition variabel
	var values = [3]int{
		90,
		85,
		75,
	}

	fmt.Println(values)

	// contoh penggunna function array
	var lagi [10]string
	fmt.Println(len(names))
	fmt.Println(len(lagi)) // akan 10 biaroun tidak isi nya, tapi dia punya panjang 10
	return "test"
}

// Tipe Data Array
/*
	1. Aray adalah tipe data yg berisikan kumpulan data dengan tipe yang sama
	2. saar membuat array, kita perlu menentukan jumlah data yang bisa di tampung oleh Array tersebut
	3. Daya Tampung array tidak bisa bertambah setelah array dibuat
*/

//Index di Array
/*
	Index		Data
	0			Eko
	1			Kurniawan
	2			Khannedy
*/

// function array
	//utnuk memanipulasi data array di golang
/*
	Operasi						Keterangan 
	len(array)					untuk mendapatkan panjang Array
	array[index]				Mendapatkan data di posisi index
	array[index] = value		Mengubah data di posisi index tertentu
*/