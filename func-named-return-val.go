package main

import "fmt"

/*
	jika tipe data yg pertama dan kedua tidak di deklarasikan
	maka tipe data terakhir wajib, dan dijadikan default
	perhatikan function di bawah
*/
func getName() (firstName, middleName string, age int16 ){
	firstName = "Arif"
	middleName = "Wardan"
	age = 20
	return
}

func main() {
	fmt.Println(getName())
	a, b, c := getName()
	fmt.Println("first name : ",a )
	fmt.Println("last name : ",b )
	fmt.Println("umur : ",c )
}
