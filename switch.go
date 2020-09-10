package main

import "fmt"

func main() {
	name:="Arif"

	switch name {
	case "Arif":
		fmt.Println("hai",name)
	case "Wardan":
		fmt.Println("hai",name)
	default:
		fmt.Println("Kenalan Dong")
	}

//	Short statement Switch

	switch length := len(name); length > 3 {
	case true:
		fmt.Println("udah benar")
	case false:
		fmt.Println("salah")
	}

//	Switch Tanpa kondisi

	switch {
	case name == "Arif":
		fmt.Println("Udah bener nih namanya")
	case name == "arip":
		fmt.Println("Salah nih")
	default:
		fmt.Println("Up gan")
	}
}
