package main

import "fmt"

func main() {
	// variabel yg nilainya ngga bisa diubah, kalo ngga dipake ngga masalah
	const name = "Aditya"

	fmt.Println(name)

	// name = "Kalo diubah gini bakal error"

	// multiple constant
	const (
		firstName = "Aditya"
		lastName  = "Ricki"
	)

	fmt.Println(firstName, lastName)
}
