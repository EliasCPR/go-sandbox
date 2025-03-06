package main

import "fmt"

// Function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello, World!")
}

// Function with parameters and no return value
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with parameters and a return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total

}

func main() {
	sayHello()
	greet("Alice")

	fmt.Println("Sum:", add(3, 4))

	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)

	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
}
