package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
}

type Contact struct{
	Email string
	Phone string
}

type Address struct{
	House string
	Area string
	State string
}

type Employee struct{
	person Person
	contact Contact
	address Address
}

func main() {
	//var void Person
	//fmt.Println(void)
	// void.Name = "Void"
	// void.Surname = "wake"
	// void.Age = 25
	//fmt.Println(void)

	// person2 := Person{
	// 	Name:    "bruce",
	// 	Surname: "wayne",
	// 	Age:     30,
	// }

	///fmt.Println(person2)

	//person3 := Person{"clark", "Kent", 28}
	//fmt.Println(person3)

	employee := Employee{
		Person{
			Name:    "Tony",
			Surname: "Stark",
			Age:     35,
		},
		Contact{
			Email: "conatct@gmail.com",
			Phone: "1234567890",
	},
		Address{
			House: "123",
			Area: "abc",
			State: "xyz",
		},
}
	fmt.Println(employee)
	fmt.Println(employee.person.Name)

}
