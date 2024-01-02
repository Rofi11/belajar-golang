package looping

import "fmt"

func Break1() string {
	for i := 0; i < 10; i++ {
		fmt.Println("perulangan", i)
		if i == 5 {
			break
		}
	}
	return "test break"
}

func Continue1() string{
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
			// akan skip semua yg genap
		}
		fmt.Println("perulangan", i)
	}
	return "test continue"
}

// Break && Continue
/*
	1. Break & Continue kata kunci yang bisa digunakan dalam perulangan
	2. Break digunakan untuk menghentikan seluruh perulangan
	3. Continue adalah digunakan unutk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjut nya
*/