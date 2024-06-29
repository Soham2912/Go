package main

import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("3rd to be executed")
	defer fmt.Println("2nd to be executed")
	fmt.Println("1st to be executed")
	
}