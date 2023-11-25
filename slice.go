package main

import "fmt"

func main() {
	// buat slice dari array
	months := [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println(months)                // [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember]
	fmt.Println(months[0:len(months)]) // [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember]
	fmt.Println(months[5:])            // [Juni Juli Agustus September Oktober November Desember]
	fmt.Println(months[:5])            // [Januari Februari Maret April Mei]
	fmt.Println(months[:])             // [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember]

	names := [...]string{"Aditya", "Ricki", "Julianto"}
	nameSlices := names[1:3]

	fmt.Println(nameSlices)
	fmt.Println(nameSlices[0])
	fmt.Println(nameSlices[1])

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu. Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days) // "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu Baru", "Minggu Baru"

	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Println(daySlice2)
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// buat slice manual (make)
	fruitSlices := make([]string, 2, 5)
	fruitSlices[0] = "Apple"
	fruitSlices[1] = "Grape"
	//fruitSlices[2] = "Mango" // error, karena capacity cuma 2, harusnya menggunakan append
	fmt.Println(fruitSlices)
	fmt.Println(len(fruitSlices))
	fmt.Println(cap(fruitSlices))

	newFruitSlices := append(fruitSlices, "Pineapple")
	fmt.Println(newFruitSlices)
	fmt.Println(len(newFruitSlices))
	fmt.Println(cap(newFruitSlices))

	newFruitSlices[0] = "Melon"
	fmt.Println(newFruitSlices)
	fmt.Println(fruitSlices)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// perbeddaan array dengan slice, array mendefinisikan jumlah data saat init, slice tidak
	//iniArray := [...]int{1, 2, 3, 4, 5}
	//iniArray := [5]int{1, 2, 3, 4, 5}
	//iniSlice := []int{1, 2, 3, 4, 5}
}
