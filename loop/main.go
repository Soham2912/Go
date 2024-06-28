package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//fmt.Println(i)

	}

	nums := []int{1, 2, 3, 4, 5}
	for index , value :=range nums {
		fmt.Println(index, value)
	}

     str:="void twt"
	 //for index,char  := range str
	 for _,char  := range str{
		 fmt.Printf("chars %c\n", char)
	 }
}