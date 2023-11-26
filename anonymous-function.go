package main

import "fmt"

type IsBlock func(string) bool

func main() {
	isBlock := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Aditya", isBlock)
	registerUser("anjing", isBlock)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}

func registerUser(name string, isBlock IsBlock) {
	if isBlock(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Hello", name)
	}
}
