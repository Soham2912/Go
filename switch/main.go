package main

import "fmt"

func main() {
	day := 6
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}


	month:="Jan"
	switch month {
	case "Jan","Feb","Dec":
		fmt.Println("Winter")
	case "Mar","Apr","May":
		fmt.Println("Spring")
	default:
		fmt.Println("other season")
}
}
