package main

import "fmt"

func main() {
	name := "Aditya"

	if name == "Aditya" {
		fmt.Println("Hello Aditya")
	} else if name == "Ricki" {
		fmt.Println("Hello Ricki")
	} else {
		fmt.Println("Siapa namamu?")
	}

	// if sort statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Hello Nama")
	}
}
