package main

import "fmt"

func main() {
	// Creating a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)
	fmt.Printf("len=%d cap=%d %v\n", len(numbers), cap(numbers), numbers)

	fmt.Printf("len=%d cap=%d %v\n", len(numbers[:0]), cap(numbers[:0]), numbers[:0])

	// Appending to a slice
	numbers = append(numbers, 6, 7)
	fmt.Println("After appending:", numbers)
	fmt.Printf("len=%d cap=%d %v\n", len(numbers), cap(numbers), numbers)

	fmt.Printf("len=%d cap=%d %v\n", len(numbers[:0]), cap(numbers[:0]), numbers[:0])

	// // Slicing a slice
	// subSlice := numbers[1:4]
	// fmt.Println("Sub-slice:", subSlice)

	// // Copying a slice
	// copiedSlice := make([]int, len(numbers))
	// copy(copiedSlice, numbers)
	// fmt.Println("Copied slice:", copiedSlice)

	// // Modifying a slice
	// numbers[0] = 10
	// fmt.Println("Modified original slice:", numbers)
	// fmt.Println("Copied slice after modification:", copiedSlice)

	// // Length and capacity of a slice
	// fmt.Println("Length of slice:", len(numbers))
	// fmt.Println("Capacity of slice:", cap(numbers))
}
