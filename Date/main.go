package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime)

	formatted := currTime.Format("02-01-2006,Monday")
	fmt.Println(formatted)
}
