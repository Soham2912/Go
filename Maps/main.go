package main

import "fmt"

func main() {
	stuRes := make(map[string]int)
	stuRes["Void"] = 49
	stuRes["Dark"] = 47
	stuRes["Gatsby"] = 45

	fmt.Println(stuRes["Gatsby"])
	fmt.Println(stuRes)

	// Checking if a key exists
	grades, exists := stuRes["David"]
	fmt.Println("Grades of David: ", grades)
	fmt.Println("Davis exists : ", exists)

	delete(stuRes, "Gatsby")
	fmt.Println(stuRes["Gatsby"])

	for index, value := range stuRes {
		fmt.Println(index, value)
	}

}
