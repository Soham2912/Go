package main

import "fmt"

	
func first(){
	fmt.Println("Hello World")
}

func add(a , b int )(int){
	return a+b
}
func main(){
   first()
   sum:= add(3,4)
   fmt.Println(sum)
}