package variabel

import "fmt"

func variabel() string{
	//buat variabel
		// cara 1
	var nama string

	nama = "Muhammad Rofi";
	fmt.Println(nama)

	// nama = "Rofi Muhammad"
	// fmt.Println(nama)

		// cara 2 ==> langsung masukan data nya, nanti golang akan pintar tau tipe data apa yang di masukan
	var Umur = 100

	fmt.Println(Umur)

		//cara 3 ==> mendecklare tanpa var diganti oleh :=
	Negara := "Indonesia";
	fmt.Println(Negara)

	// utk perubahan isi data variabel
	Negara = "Malaysia"

	// Deklarasi Multiple Variable
		// di golang bisa langsung buat banyak variabel & code yang dibat lebih bagus
	var (
		firstName = "Rofi"
		lastName = "Muhammad"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	return "test variabel"
}
// Variabel
	// Tempat untuk menyimpan data
	// digunakan agar kita bisa mengakses data yang sama dimanapun  kita mau
	// di golang kita hanya bisa menyimpan tipe data yang saam, jika ingin menyimpan data yang berbeda-beda jenis, kita harus membaut beberapa variabel
	// untuk membuat variable, kita bisa menggunakan kata kunci var, lalu dengan nama variable dan tipe datanya

// kata kunci var
	// di golang kata kunci var saat membuat variable tidaklah wajib
	// asalkan saat membuat variable kita langsung menginilisasikan data nya
	// agar tidak perlu menggunakan lata kunci Var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut.