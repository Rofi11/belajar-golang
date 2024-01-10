package function

import (
	"fmt"
)

// contoh Deffer fucntion
func logging() {
	fmt.Println("Selesai Memanggil function")
}

func RunApplicationDefer(value int) int{
	// call defer function
	  // pakai defer biasakan di atas
	defer logging()
	fmt.Println("Run Application")

	result := 10/value
	return result
}

func endApp(){
	fmt.Println("End App")
}

func RunAppPanic(error bool) string{
	defer endApp()
	if error {
		panic("APPLICATION ERROR")
	}
	return "APPLICATION SUCCESS"
}

func endAppRecover(){
	// fmt.Println("End App")
	message := recover()
	if message != nil{
		fmt.Println("Tejadi kesalahan dengan message", message)
	}
	fmt.Println("Apllikasi selesai")
}


func RunAppRecover(error bool) string{
	defer endAppRecover() // melakukan proses recover panic nya di beda fucntion dalam defer
	if error {
		panic("APPLICATION ERROR")
	}
	fmt.Println("Apllication berjalan")
	return "test Recover"
}

// Defer
	// akan tetap dijalan kan biarpun error, simpan paling atas
/*
	1. Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
	2. Defer function akan selalu di ekseskusi walaupun terjadi error di function yang di eksekusi
*/

// Panic
	// akan menghentikan program jika error
/*
	1. panic function adalah fucntion yang bisa kita gunakan untuk menghentikan program
	2. Panic function biasanya di panggil ketika terjadi error pada saat program kita berjalan
	3. Saat panic fucntion di panggil, program akan terhenti, namun defer fucntion tetap akan di eksekusi
*/

// Revocer
	// untuk menangkap data panic, dan kode uakan terus berjalan
/*
	1. Recover adalah fucntion yang bisa kita gunakan untuk menangkap data panic
	2. Dengan recover proses panis akan nterhenti, sehingga program berjalan
*/