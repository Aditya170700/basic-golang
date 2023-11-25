package main

import "fmt"

func main() {
	goodbye := getGoodBye

	fmt.Println(goodbye("Aditya"))
}

func getGoodBye(name string) string {
	return "Good bye " + name
}
