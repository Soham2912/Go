package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 42
	fmt.Printf("type %T \n", num)

	var data float32 = float32(num)
	fmt.Printf("type %T \n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Printf("type %T \n", str)

	str = "1234"
	num, _ = strconv.Atoi(str)
	fmt.Printf("type %T \n", num)

	var num_float float64

	str = "3.14"
	num_float, _ = strconv.ParseFloat(str, 64)
	fmt.Printf("type %T ", num_float)
}
