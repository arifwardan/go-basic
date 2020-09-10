package main

import "fmt"

func main() {
	var array = [5] int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(array)
	fmt.Println(len(array)) // untuk menghitung panjang dari array
	array [0] = 100
	fmt.Println(len(array))
}
