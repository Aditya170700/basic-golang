package main

import "fmt"

func main() {
	// buat ulang tipe data baru dari tipe data yg udah ada
	// biasa digunakan untuk membuat alias dari tipe data yg udah ada

	type NoKTP string // tipe data baru (NoKTP) dari tipe data string

	var ktpAditya NoKTP = "9128798127982113"
	fmt.Println(ktpAditya)
	fmt.Println(NoKTP("7162798127982113")) // konversi string ke NoKTP
}
