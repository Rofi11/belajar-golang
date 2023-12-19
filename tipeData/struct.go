package tipeData

// mendifinisikan object/ struct di golang
type Person struct{
	Name string
	Age int
}

// function yang mengembalikan object/struct person
func Struct(name string, age int) Person{
	person := Person{
		Name : name,
		Age : age,
	}

	return person
} 