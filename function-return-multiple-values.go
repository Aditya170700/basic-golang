package main

import "fmt"

func main() {
	fName, lName := getFullName()
	fmt.Println(fName, lName)

	fName, _ = getFullName()
	fmt.Println(fName)
}

func getFullName() (string, string) {
	return "Aditya", "Ricki"
}
