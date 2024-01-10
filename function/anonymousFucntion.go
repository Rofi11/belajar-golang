package function

type BlackList func(string) bool

func registeruser(name string, blacklist BlackList) string {
	var hasil string
	if blacklist(name) {
		hasil = "you are blocked" + " " + name
		return hasil
	} else {
		hasil = "welcome" + " " + name
		return hasil
	}
}

func AnonymousFunction() string {
	// anonymous function
	blackList := func(name string) bool {
		return name == "admin"
	}
	finalHasil := registeruser("admin", blackList)
	return finalHasil
}

// Anonymous Function
/*
	1. sebelumnya setiap membuat function, kita akan selalu memberikan sebauh nama pada function tersebut
	2. namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
	3. hal tersebut dinamakan anonymous function, atau function tanpa nama
*/