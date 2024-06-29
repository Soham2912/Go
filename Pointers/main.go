package main

import "fmt"

func modifyByReference(num *int) {
	*num = *num * 2
}

func main() {
	// var num int
	num := "void"

	// var ptr *int

	ptr := &num

	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	value := 10
	modifyByReference(&value)
	fmt.Println(value)

}
