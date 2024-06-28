package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "lol"
	name[1] = "kek"

	fmt.Println("Names", name)
	var numbers = [5]int{1,2,3,4,5}
	fmt.Println("Numbers", numbers)
	fmt.Println("Length of numbers", len(numbers))
	fmt.Println("Length of name", len(name))
	

}