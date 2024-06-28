package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	var void Person
	fmt.Println(void)
	void.Name= "Void"
	void.Surname= "wake"
	void.Age= 25
	fmt.Println(void)

	person2 := Person{
		Name: "bruce",
		Surname: "wayne",
		Age: 30,
	}

	fmt.Println(person2)

}
