package looping

import "fmt"

func ForLoops() string {
	counter := 1

	for counter <= 10 { // kodisi
		fmt.Println("perulangan ke ", counter)
		counter++ // invrement/ decrement
	}

	return "test perulangan"
}

// for dengan statement
func ForLoopsWithStatement()string{
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	return "test for loop dengan statement"
}

// for range
func ForRange()string{
	slice := []string{"Eko", "Kurniawan", "Khannedy"}
	
	// ini cara manual
	// for i := 0; i < len(slice); i++{
	// 	fmt.Println(slice[i])
	// }

	// ini cara pakai for range
	// for i, value := range slice{
	// 	fmt.Println("index", i, "=", value)
	// }

	//jika data index nya tidak di butuhkan di ganti oleh undercore(_)
	for _, value := range slice{
		// jika pakai underscore memberitahu golang variable index tidak di butuhkan
		fmt.Println(value)
	}

	//perulangan untuk map
		// index nya di ganti key jika untuk map
	person := make(map[string]string)
	person["name"] = "Rofi"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
	return "test For Loop Range"
}


// di golang hanya ada satu perulangan nyaitu for

// For
/*
	1. Dalam Bahasa pemograman. Biasanya ada fitur yang bernama perulangan
	2. salah satu fitur perulangan adalah for loops
*/

// for dengan statement
/*
	> dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa di tambahkan di for :
	1. Init statement sebelum fo di eksekusi
	2. post statement, nyaitu statement yang akan selalu di eksekusi di akhit tiap perulangan
*/

// for Range
/*
	1. For bisa di gunakan untuk melakukan iterasi terhadap semua dara collection
	2. Data collection contoh nya array, slice, map
*/
