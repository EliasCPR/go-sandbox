package main

import "fmt"

func main() {
	// Arrays
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println("Array:", arr)

	for i, v := range arr {
		fmt.Printf("the value at index %d is %d\n", i, v)
	}

	// Modifying the array
	arr[0] = 10
	arr[1] = 20
	fmt.Println("Modified Array:", arr)

	// Creating a new array
	var newArr [3]int
	newArr[0] = 7
	newArr[1] = 8
	newArr[2] = 9
	fmt.Println("New Array:", newArr)

	// Copying an array
	var copiedArr [5]int
	for i := range arr {
		copiedArr[i] = arr[i]
	}
	fmt.Println("Copied Array:", copiedArr)
}
