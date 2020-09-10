package main

import "fmt"

func main() {
	type nama string
	type Married = bool

	var firstName nama = "Arif Wardan"
	var statusMarried Married = false

	fmt.Println(firstName)
	fmt.Println("Status Pernikahan = ",statusMarried)

}
