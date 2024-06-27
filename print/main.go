package main

import "fmt"

func main() {
	age := 25
	name := "Void"
	height := 180.00

	fmt.Println("age", age, "Height", height, "name", name)
	fmt.Printf("height %.2f\n", height)
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of name is %T\n", name)
}
