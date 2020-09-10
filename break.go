package main

import "fmt"

func main() {
	for i := 1; i<= 10; i++ {
		if i >=5 {
			break
		}
		fmt.Println("perulangan ke ",i)
	}

	for j := 1; j<= 10; j++ {
		if j <=5 {
			continue
		}
		fmt.Println("perulangan ke ",j) //akan di skip jika if di atas bernilai false
	}
}
