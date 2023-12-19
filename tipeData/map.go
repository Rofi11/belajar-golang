package tipeData

import "fmt"

func Map() {
	// tipe data map

	// membuat dengan ada nya data (langsung declarayion key : value)
	person := map[string]string{
		"name" : "Rofi",
		"address" : "Subang",
	}

	fmt.Println(person["name"]) // panggil berdasarkan type key nya
	// menambah
	person["title"] = "Programmer"
	fmt.Println(person["title"])
	// merubah
	person["address"] = "Permata"
	fmt.Println(person["address"])

	// membuat dari awal (data nya kosongan)
	book := make(map[string]string)
	book["title"] = "Buku Go-lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "ups"
	fmt.Println(book)

	//delete map
	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))// lenght nya berdasarkan jumlah data di dalam nya
}

// Tipe data Map
/*
	Tipe Data Map
	1. pada Array atau slice, untuk mengakses data, kita mneggunakan index Number dimulai dari 0
	2. Map adalah tipe data lain yang berisikan kumpulan data yang sama, namu kita bisa menntukan jenis tipe data index yang akan kita gunakan
	3. sederhananya, Map adalah tipe data kumpulan key-value (kata lunci - nilai), dimana kata kuncinya bersifat unik, tidak bolhe sama
	4. berbeda dengan array dan slice, jumlha data yangkita masukkan ke dalam Map boleh sebanyak banyak , asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka akan secara otomatis data sebelumnya akan di ganti dengan data baru

	person := map[string]string
				type key, type value nya
*/


// function map
/*
	Operasi							keterangan
	len(map)						Untuk mendapatkan jumlah data di map (tergantung jumlah data di map nya)
	map(key)						Mengambil data di map dengan key
	map(key) = value				Mengubah data di map dengan key
	make(map[TypeKey]TypeValue)		Membuat map baru
	delete(map, key)				Menghapus data di map dengan key
*/