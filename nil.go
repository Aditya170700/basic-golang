package main

import "fmt"

func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}
