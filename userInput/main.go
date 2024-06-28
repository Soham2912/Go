package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input your name:")
	var name string
	//fmt.Scan(&name)
	//fmt.Println("Hello", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString(' ')
	fmt.Printf("Hello %s", name)

}
