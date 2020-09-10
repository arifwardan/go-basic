package main

import "fmt"

func main() {
	name:= "arif"
	if name != "arif" {
		fmt.Println(name)
	}else {
		fmt.Println("hello", name)
	}

//	if Short Statement Cuma ada di golang
	if length:=len(name); length > 3 {
		fmt.Println("panjang nama ",name," = ", length)
	}

}
