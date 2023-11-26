package main

import "fmt"

type Location struct {
	City, Province, Country string
}

func main() {
	location := Location{"Gunungkidul", "DIY", "Indonesia"}

	fmt.Println(location)

	// ubah city ke "Jogja"
	ChangeCityWithoutPointer("Jogja", location)
	// city tidak berubah
	fmt.Println(location)

	// ubah city ke "Sleman"
	ChangeCityWithPointer("Sleman", &location)
	// city berubah, karena menggunakan pointer
	fmt.Println(location)
}

func ChangeCityWithoutPointer(city string, location Location) {
	location.City = city
}

func ChangeCityWithPointer(city string, location *Location) {
	location.City = city
}
