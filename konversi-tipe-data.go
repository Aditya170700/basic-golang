package main

import "fmt"

func main() {
	// number
	// perlu diperhatikan kapasitasnya kalo konversi dari tipe data yg kapasitasnya besar ke yang lebih kecil
	// ngga error cuma jadi rancu aja nilainya
	var nilai32 int32 = 8127
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// string
	var name = "Aditya"
	var a = name[0]         // ini dapetnya uint8
	var aString = string(a) // konversi dari uint8 ke string

	fmt.Println(name)
	fmt.Println(a)
	fmt.Println(aString)
}
