package main

import "fmt"

func getFullName () (string, string){
	return "Arif", "Wardan"
}

func main() {
	_, lastName := getFullName()
	// kalau kita tidak peduli dengan salah satu return valuenya gunakan _ (garis bawah)
	//fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(getFullName())
}
