package main

import "fmt"

func main() {
//	tipe data Map Hampir sama dengan array associative di php
//	var person map[string]string = map[string]string{} ini deklarasi panjang
	person := map[string]string{
		"name": "Arif", //harus petik " dua berbeda dengan php
		"age" : "20",
		/*
			bisa ditambah datanya asalkan dengan key unik, jika yang digunakan
			key yang sudah ada maka yang baru akan menimpa yang lama
		*/
	}

	person["title"] = "Programmer"
	person["division"] = "Backend"
	person["name"] = "Wardan"
	person["ktp"] = "asdf"

	fmt.Println(person["ktp"])
	delete(person, "ktp")
	fmt.Println(person)
	fmt.Println(person["name"])

	book := make(map[string]string)
	book["name"] = "Belajar Go"
	book["Author"] = "Arif Wardan"

	fmt.Println(book)


}
