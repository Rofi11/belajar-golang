package function

func penjumlahan(data1 int64, data2 int64) int64 {
	hasil := data1 + data2

	return hasil
}

func perkalian(data1 int64, data2 int64) int64 {
	hasil := data1 * data2

	return hasil
}

func pengurangan(data1 int64, data2 int64) int64 {
	hasil := data1 - data2

	return hasil
}

func pembagian(data1 int64, data2 int64) int64 {
	hasil := data1 / data2

	return hasil
}

func FunctionReturnValue(data1 int64, data2 int64, operasi string) int64 {
	var finalHasil int64
	switch operasi {
	case "penjumlahan":
		finalHasil = penjumlahan(data1, data2)
	case "pengurangan":
		finalHasil = pengurangan(data1, data2)
	case "perkalian":
		finalHasil = perkalian(data1, data2)
	case "pembagian":
		finalHasil = pembagian(data1, data2)
	default:
		finalHasil = penjumlahan(data1, data2)
	}
	return finalHasil
}

// function return value
/*
	1. function bisa mengembalikan data
	2. Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut
	3. jika function tersebut kita deklarasikan dengan tipe data pengemballian, maka wajib di dalam function nya kita harus mengembalikan data
	4. Untuk mengembalikan data dari function , kita bisa menggunakan kata kunci return, diikuti dengan data nya
*/