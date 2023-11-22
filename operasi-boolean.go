package main

import "fmt"

func main() {
	a := true
	b := false

	fmt.Println(a && b) // nilainya true kalo semua true
	fmt.Println(a || b) // nilainya true kalo salah satu true
	fmt.Println(!a)     // kebalikan dari a

	nilaiAkhir := 90
	absensi := 80
	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi > 80
	lulus := lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}
