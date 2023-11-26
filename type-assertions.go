package main

import "fmt"

func main() {
	var result interface{} = random()

	// manual ngga error
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	// manual error. ini akan panic, karena result outputnya string
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

func random() interface{} {
	return "OK" // string
}
