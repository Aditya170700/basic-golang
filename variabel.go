package main

import "fmt"

func main() {
	// cara 1 - init tanpa value dengan tipe data
	var name1 string
	name1 = "Aditya"
	fmt.Println(name1)

	// cara 2 - init value tanya tipe data
	var name2 = "Aditya"
	fmt.Println(name2)

	// cara 3 -  init pakai :=
	name3 := "Aditya"
	fmt.Println(name3)
	name3 = "Ricki"
	fmt.Println(name3)

	// cara 4 - init multiple variabel
	var (
		firstName = "Aditya"
		lastName  = "Ricki"
	)
	fmt.Println(firstName, lastName)
}
