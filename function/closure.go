package function

import "fmt"

func CLosure() string {
	// global scope function Closure
	counter := 0
	name := "Eko"

	increment := func(){
		// block scope function increment
		name := "Budi"
		fmt.Println("Increment")
		fmt.Println(name)
		counter++
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)
	return "test penggunaan closure"
}

// Closure
/*
	1. Closure adalah kemampuan sebuah fucntion berinteraksi dengan data-data disekitarnya dalam scope yang sama
		> scope lingkup kerja sebuah variabel
	2. harap gunakan fitru closure ini dengan bijak saat kita membuat applikasi
*/