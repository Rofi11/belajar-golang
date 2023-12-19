package tipeData

import "fmt"

func slice() {
	var months = [...]string{
		"Januari",
		"februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// slice
	// var slice1 = months[4:7]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// membuat slice baru
	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "rofi") // membuat array baru karena pas append sudah melewati limit array nya
	// fmt.Println((slice3))
	slice3[1] = "Bukan Desember" // tidak akan mempengaruhi slice 2 karena slice 3 menjadi array baru tersendiri
	fmt.Println(slice3)
	fmt.Println(months)

	// Membuat Array baru dari awal
	newSlice := make([]string, 2, 5) // akan dibatkan array baru yang type string, len 2, cap 5
	fmt.Println(newSlice) //1. array kosong
	newSlice[0] = "Rofi" //2. masukan data
	newSlice[1] = "Muhammad"
	fmt.Println(newSlice) //3. sudah akan ada data nya
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))


	// mengcopy slice
	copySlice := make([]string, len(newSlice), cap(newSlice)) // pastikan len dan cap nya sama dengan yg mau di copy, agat tidak terpotong isi nya
	copy(copySlice, newSlice)
	fmt.Println(copySlice) // isinya akan seperti newSlice


	//noted perbedaan array dan slice
	// hati hati dalam membuat array dan slice
	// array
	iniArray := [...]int{1,2,3,4,5}
	fmt.Println(iniArray)
	// slice
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniSlice)
}

// Tipe Data Slice
/*
	1. Tipe data slice adalah potongan dari data Array
	2. Slice mirip dengan Array, yang menbedakan adalah ukuran Slice bisa berubah
	3. slice dan array selalu tekoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di array
*/

// Detail Tipe Data Slice
/*
	1. Tipe Data Slice memiliki 3 Data, yaitu pointer, length dan capcity
	2. pointer adalah penunjuk data pertama di array para slice
	3. length adalah panjang dari slice
	4. capacity adalah kapasitas dari slice, dimana length tidak boleh dari capacity
*/

// Membuat Slice dari Array
/*
	jika isi array 10

	Membuat Slice				keterangan
	1. array[low:high]			Membuat slice dari array dimulai index low sampai index sebelum high
	ex : 2-9, karena di ambil sebelum high, data terakhir tidak akan di ambil
	2. array[low:]				Membuat slice dari array dimulai index low sampai indes akhir di array
	ex : 4-9, di ambil sampai akhir
	3. array[:high]				Membuat slice dari array dimulai dari index 0 sampai index sebelum high
	ex : dimulai dari 0-5, karena sebelum high
	4. array[:]					Membuat slice dari array dimulai index 0 sampai index akhir di array
	ex: 0-9, smeua data array dari index ke 0 dampai akhir
*/

// noted
/*
	1. jika belum tau isi array pakai [...], titik 3 kali kaya spread operator jika di javascript
	2. mengubah isi slice akan berdampak perubahan ke array nya juga
	3. membuat array baru dari awal pakai make
*/


// Function Slice
/*
	Operasi										Keterangan
	len(slice)									Untuk mendapatkan panjang
	cap(slice)									untuk mendapatkan kapasitas
	appends(slice, data)						Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
	make([]typeData, length, capacity)			Membuat slice baru
	copy(destination, source)					Menyalin slice dari source ke destination
*/

// noted
/*
	1. function make
		> Membuat array baru tanpa mempengaruhi array lain
		> karena tidak di masukan ke dalam variabel nya
		> fungsi nya seperti membuat useState arrat kosong
	2. function copy
		> Mengcopy array slice yang sudah ada
*/