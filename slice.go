package main

import "fmt"

func main() {
	var month = [...] string{ // [...] digunakan jika kita tidak tau mau buat arraynya berapa
		"januari",
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

	fmt.Println(len(month))
	fmt.Println(cap(month))

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	month[5] = "ganti";
	fmt.Println(slice1)

	slice1[0] = "ubah"
	fmt.Println(month)
//	ketika array di ubah atau slice di ubah maka yg lainya juga keubah
	var slice2 = append(slice1, "Arif")
	/*
		ketika menambahkan array baru dengan append() dan kapasitas dari array sudah penuh maka dia
		akan membuat array baru, yg otomatis jika data dari slice2 kita ubah maka array month
		tidak akan berubah, tapi jika kapasitas dari array masih ada dia tidak akan membuat array baru yg
		otomatis jika kita mgubah slice2 maka array month juga akan ikut berubah
	*/
	fmt.Println(slice2)
	fmt.Println(month)

	makeSlice := make([]string, 2, 2)
	makeSlice[0] = "Arif"
	makeSlice[1] = "Wardan"

	fmt.Println(makeSlice)

	copySlice := make([]string, len(makeSlice), cap(makeSlice))
	copy(copySlice, makeSlice)
	copySlice[1] = "Aja"
	fmt.Println(copySlice)
	fmt.Println(makeSlice)

	iniArray := [...]int{1,2,3,4,5,6,}
	iniJugaArray := [6]int{1,2,3,4,5,6,}
	iniSlice := []int{1,2,3,4,5,6,}

	fmt.Println(iniArray)
	fmt.Println(iniJugaArray)
	fmt.Println(iniSlice)

}
