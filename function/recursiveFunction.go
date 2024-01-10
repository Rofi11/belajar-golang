package function

func RecursiveFunction(value int) int {
	if value == 1 {
		return 1 // utk kondisi berhenti nya, ketika satu balikan 1, jadi pas satu akan berhenti, agar tidka jadi infinete loop
	} else {
		return value * RecursiveFunction(value-1)
		// 5 * 4 * 3 * 2 *1
	}
}

func FactorialLoop(value int) int {
	// cara normal pakai loop
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	// 5 * 4 * 3 * 2 *1

	return result
}

// Recursive Fucntion
/*
	1. Recursive function adalah function yang memanggil function dirinya sendiri
	2. kadang dalam pekerjeaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive fucntion
	3. contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah factorial
		> factorial adalah perkalian dengan angka angka sebelum nya
*/
