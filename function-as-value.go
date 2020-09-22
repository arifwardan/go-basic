package main

import "fmt"

func getGreeting(name string)string{
	return "Hello " + name
}

func main() {
	sayHello := getGreeting
	result := sayHello("arif")
	fmt.Println(result)
}
