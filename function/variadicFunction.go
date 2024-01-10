package function

func VariadicFunction(numbers ...int64) int64 {
	total := 0
	for _, value := range numbers {
		total += int(value)
	}
	return int64(total)
}

func VariadicFunctionSliceParameter(numbers ...int) int64 {
	total := 0
	for _, value := range numbers {
		total += int(value)
	}
	return int64(total)
}

// Variadic Function
/*
	1. Parameter yang berbeda di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs(variabel arguments) (Wajib di posisi terakhir)
	2. Varags (variabel arguments) artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam array
	3. apa bedanya dengan parameter biasa dengan tipe data array?
		> jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
		> jika parameter mengggunakan varargs (variabel arguments), kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma
	4. function yg memilki varags (variable arguments ) di sebut variadic function
*/

// bagaimana jika punya variabel slice dijadikan parameter
// Slice Parameter
/*
	1. kadang ada kasus dimana kita menggunakan Variadic function, namun memiliki vatiable berupa slice
	2. kita bisa menjadikan slice sebagai varargs(variabel arguments) paramenter
*/