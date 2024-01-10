package function

func getGoodBye(name string) string {
	return "GoodBye" + " " + name
}

func FunctionValue() string {
	goodBye := getGoodBye("zika")
	return goodBye
}

// funtion adalah first calss citizen
/*
	1. Fcuntion adalah ifrst class citizen
	2. Function juga meruapakn tipe data,  dan  bisa disimpan didalam variabel
*/