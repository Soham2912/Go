package main

import "fmt"

func first() {
	fmt.Println("Hello World")
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func mul(a , b int) int {
	return a * b
}

func main() {
	first()
	sum := add(3, 4)
	fmt.Println(sum)
	fmt.Println(mul(3, 4))	
}
