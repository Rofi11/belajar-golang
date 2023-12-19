package variabel

import "fmt"

func constantVariabel()string {
	const (
		firstName string = "Rofi"
		lastName = "Muhammad"
		value = 1000
	)

	// jika variabel constant boleh utk tidak di gunakan, tapi kalo var wajib di gunakan 

	fmt.Println(firstName, lastName)

	return "test constant"
}