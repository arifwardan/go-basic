package main

import "fmt"

func main() {
	var a = 20
	var b = 2
	var c = a + b
	var d = c - b
	var e = (a + c) / b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
//	Augmented Assignments
	var i = 1
	i += 10 //=>  i = i + 10
	fmt.Println(i)
//	Unary Operator
	i++
	fmt.Println(i)
	var positif = +100
	var negativ = -100
	fmt.Println(positif)
	fmt.Println(negativ)
}
