package function

import "fmt"


func sayHello(){
	fmt.Println("sayHello")
}

func TestFunction() string {
	sayHello()

	return "test function"
}