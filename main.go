package main

import (
	"fmt"

	// "example/packages/conditional"
	// "example/packages/tipeData"
	// "example/packages/looping"
	// "example/packages/function"
	// "example/packages/structt"
	// "example/packages/interfacet"
	// "example/packages/nill"
	"example/packages/pointer"
)

// call semua task ke sini
func main(){
	// Tipe data array
	// fmt.Println(tipeData.Array())
	// fmt.Println(tipeData.String("toni"))

	//tipe data object/struct
	// fmt.Println(tipeData.Struct("Rofi", 26))

	// condtional
	// fmt.Println(conditional.SwitchExpression())
	
	// fmt.Println(conditional.SwitchExpressionShort())
	// fmt.Println(conditional.SwitchTanpaKondisi())
	
	// Looping
	// fmt.Println(looping.ForLoops())
	// fmt.Println(looping.ForLoopsWithStatement())
	// fmt.Println(looping.ForRange())

	// break and continue
	// fmt.Println(looping.Break1())
	// fmt.Println(looping.Continue1())

	//function
	// fmt.Println(function.TestFunction())
	// fmt.Println(function.FunctionParameter("Muhammad", "Rofi"))
	// fmt.Println(function.FunctionReturnValue(10, 5, "perkalian"))
	// fmt.Println(function.FunctionReturnMultipleValue("Eko", "Khannedy"))
	// fmt.Println(function.NamedReturnValues())
	
	// case jika bentuk nya biasa 
	// fmt.Println(function.VariadicFunction(10,10,10,10,10))
	//case jika kita sudah punya duluan variabel slice
	// numbers := []int{10,10,10,10,10,20}
	// fmt.Println(function.VariadicFunctionSliceParameter(numbers...)) //muaskan data nya harus ada ..., titik 3 nya

	// fmt.Println(function.FunctionValue())
	// fmt.Println(function.FucntionAsParameter())
	// fmt.Println(function.AnonymousFunction())


	// fmt.Println(function.FactorialLoop(5))
	// fmt.Println(function.RecursiveFunction(5))
	// fmt.Println(function.CLosure())

	// defet-panic-recover fucntion (ERROR HANDLING FUNCTION)
	// fmt.Println(function.RunApplicationDefer(1))
	// fmt.Println(function.RunAppPanic(true))
	// fmt.Println(function.RunAppRecover(true))


	//struct
	// fmt.Println(structt.Struct())
	// fmt.Println(structt.VariabelStruct())
	// fmt.Println(structt.StructLiterals())
	// dataStruct := structt.Struct()
	// fmt.Println(structt.StructMethod(dataStruct, "Eko"))

	//Interface
	// dataStructInterface := interfacet.PersonStructInterface()
	// fmt.Println(interfacet.Interfacet(dataStructInterface))
	// fmt.Println(interfacet.InterfaceKosong(3))
	// fmt.Println(interfacet.TypeAssertion())
	// fmt.Println(interfacet.TypeAssertionsWithSwitch())

	// nil
	// var person map[string]string = nill.Nill("Rofi")
	// if person == nil {
	// 	fmt.Println("Data kosong")
	// } else {
	// 	fmt.Println(person)
	// }

	//pointer
	// fmt.Println(pointer.CobaPassByValue())
	// fmt.Println(pointer.CobaPointer())
	// fmt.Println(pointer.OperatorBintang())
	// fmt.Println(pointer.CobaOperatorNew())
	// fmt.Println(pointer.PointerDiFunction())
	fmt.Println(pointer.PointerDiMethod())
}