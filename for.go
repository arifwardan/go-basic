package main

import "fmt"

func main() {
	for i:=1; i <= 10; i++{
		for j:=1; j <=10; j++ {
			if(i < j){
				fmt.Println("*")
			}else {
				fmt.Println("-")
			}
		}
		fmt.Println("\n")
	}

	counter := 1
	for counter <=10 {
		fmt.Println("counter ", counter)
		counter++
	}

//	for range
//	slice
	persons := []string{"Arif wardan", "20", "Riau", "Programmer"}
	for _, person := range persons{
		fmt.Println("Data", person)
	}
	/*
		================================!NOTE!=================================
		Gunakan _(underscore) jika ada sebuah variabel yg tidak ingin di gunakan
		================================!NOTE!=================================
	*/
//	map / array assoc di php
	value := map[string]string{"name":"Arif wardan","age":"20","adress":"Riau","division":"Programmer"}
	for key, val := range value{
		fmt.Println("Key", key, "Value", val)
	}

}
