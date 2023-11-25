package main

import "fmt"

func main() {
	name := "Aditya"

	switch name {
	case "Aditya":
		fmt.Println("Hello Aditya")
	case "Ricki":
		fmt.Println("Hello Ricki")
	default:
		fmt.Println("Siapa namamu?")
	}

	// sort statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Hello Nama")
	}

	// tanpa kondisi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Hello Nama")
	}
}
