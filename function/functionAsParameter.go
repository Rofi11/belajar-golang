package function

// type Declaraion utk mempermudah penulisan yg panjang  menjasi alias
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	// function parameter nya di pakai di sini
	nameFiltered := filter(name)
	return nameFiltered
}

// function utk jadi parameter utk filtering
func spamFilter(name string) string {
	if name == "Anjing" {
		return "Jangan Pakai Bahasa Kasar"
	} else {
		return name
	}
}

// pakai di sini untk call function dan pakai parameter nya
func FucntionAsParameter() string {
	filter := sayHelloWithFilter("Anjing", spamFilter)
	return filter
}

// function as parameter
/*
	1. function tidak hanya bisa kita simpaan didalam variable sebagai value
	2. namun juga bisa kita gunakan sebagai parameter untuk function lain
*/

// fucntion type Declarations
/*
	1. kadang jika fcuntion terlalu panjang, agak ribet untuk menuliskan didalam parameter
	2. Type Declaraions juga bisa dihunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter
*/