package main

import "fmt"

func div(a, b float32) (float32, error) {
	//return a / b
	if b==0 {
		return 0 , fmt.Errorf("cannot divide by 0")
	}
	return a/b , nil
}

func main() {
	ans, err := div(10,0)
	fmt.Println(ans,err)
}