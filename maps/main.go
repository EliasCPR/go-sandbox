package main

import "fmt"

func main() {
	// Create a map
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["Alice"] = 25
	myMap["Bob"] = 30
	myMap["Charlie"] = 35

	// Access a value by its key
	fmt.Println("Alice's age:", myMap["Alice"])

	// Check if a key exists
	age, exists := myMap["Bob"]
	if exists {
		fmt.Println("Bob's age:", age)
	} else {
		fmt.Println("Bob not found")
	}

	// Delete a key-value pair
	delete(myMap, "Charlie")

	// Iterate over the map
	for key, value := range myMap {
		fmt.Printf("%s is %d years old\n", key, value)
	}
}
