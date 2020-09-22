package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(23,32,34,5,34,567,8,)
	fmt.Println(total)

	slice := []int{23,54,76,34,76}
	total = sumAll(slice...)
	fmt.Println(total)
}
