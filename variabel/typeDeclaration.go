package variabel

import "fmt"

func TypeDeclaraion() string {

	// ini NoKTP utk alias string
		// type string dijadikan NoKTp
	type NoKTP string
	type Married bool // boolean

	var ktpRofi NoKTP = "111111"
	var marriedStatus Married = true

	fmt.Println(ktpRofi)
	fmt.Println(NoKTP("222222"))
	fmt.Println(marriedStatus)

	return "test type Declarion"
}