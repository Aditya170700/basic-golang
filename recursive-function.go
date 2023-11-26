package main

import "fmt"

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecusrive(10))
}

func factorialRecusrive(value int) int {
	if value == 1 {
		return 1
	}

	return value * factorialRecusrive(value-1)
}

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}
