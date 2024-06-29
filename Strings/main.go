package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str:="1 2 3 4 5 1 1 6"
	cnt :=strings.Count(str, "1")
	fmt.Println(cnt)

	str="                               void    "
	str=strings.TrimSpace(str)
	fmt.Println(str)

	str="apple"
	str=strings.Repeat(str, 5)
	fmt.Println(str)

	str1:="apple"
	str2:="seeds"
	result:=strings.Join([]string{str1,"lol",str2},",")
	fmt.Println(result)
}
