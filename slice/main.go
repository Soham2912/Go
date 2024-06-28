package main

import "fmt"

func main() {
	// slice is a reference type
	// slice is a dynamic array
	// slice is a view of an array
	// slice is a reference to an array
	// slice is a reference to a contiguous segment of an array

	numbers := []int{1, 2, 3, 4, 5}
	numbers= append(numbers, 6,6,78,8,9)
	fmt.Println("numbers", numbers)
	fmt.Println("Length",len(numbers))
	fmt.Println("Capacity", cap(numbers))

	nums:=make([]int, 3 ,5)
	fmt.Println("nums", nums)
	nums = append(nums, 6,6)
	fmt.Println("nums", nums)
	fmt.Println("Length",len(nums))
	fmt.Println("Capacity", cap(nums))
	nums = append(nums, 9)
	fmt.Println("nums", nums)
	fmt.Println("Length",len(nums))
	fmt.Println("Capacity", cap(nums))


}