package tipeData

import "fmt"

func test(){
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)
}


// Tipe Data Number di golongkan dalam 2
	// > Interger
	// > Floating Point

// Dalam golang ada Nilai spesifik yang bisa di masukan(limit yang bidsa di masukan sesuai kebutuhan) untuk interger dan floating Point
	// INTERGER
	// tipe int ==> memiliki angka negatif
	// Tipe data || Nilai Minimum || Nilai Maksimum
	// int8		|| -128			|| 127
	// int16	|| -32768		|| 32767
	// int32	|| -2147483648	|| 2147483647
	// int64	|| -9223372036854775808 || 9223372036854775807

	// tipe uint (unsignInterger) ==> tidak memiliki angka negarif dan nilai maksimum nya lebih tinggi dari int
	// Tipe data || Nilai Minimum || Nilai Maksimum
	// uint8	 || 0			  || 255
	// uint16	 || 0			  || 65535
	// uint32	 || 0			  || 4294967295
	// uint64	 || 0			  || 18446744073709551615


	//Floating Point
	// Tipe data || Nilai Minimum || Nilai Maksimum
	// float32	 || 1.18 x 10 pangkat minus 38 || 3.4 x 10 pangkat 38
	// float64	 || 2.23 x 10 pangkat minus 308 || 1.80 x 10 pangkat 308 
	// complex64 || complex numbers with float32 real and imaginary parts.
	// complex128 || complex numbers with float64 real and imaginary parts.


	// Alias ===> semacam nama lain utk data yang sudah ada
	// Tipe Data		||	Alias Untuk
	// byte				||	uint8
	// rune				||	int32
	// int				|| 	Minimal int32 ==> sesuai dengan sistem opersasi jiga 64byte maka jadi int64
	// uint				|| 	Minimal uint32 ==> sesuai dengan sistem opersasi jiga 64byte maka jadi int64




