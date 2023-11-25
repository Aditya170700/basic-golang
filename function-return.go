package main

import "fmt"

func main() {
	fmt.Println(getHello("Aditya"))
}

func getHello(name string) string {
	return "Hello " + name
}
