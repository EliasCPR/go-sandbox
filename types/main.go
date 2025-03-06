package main

import "fmt"

func main() {
	// Boolean type
	var b bool = true
	fmt.Println(b)

	// Numeric types
	var i int = 42
	var f float64 = 3.14
	var c complex128 = complex(1, 2)
	fmt.Println(i, f, c)

	// String type
	var s string = "Hello, Go!"
	fmt.Println(s)

	// Array type
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// Slice type
	var sl []int = []int{1, 2, 3, 4, 5}
	fmt.Println(sl)

	// Map type
	var m map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Println(m)

	// Struct type
	type Person struct {
		Name string
		Age  int
	}
	var p Person = Person{Name: "Alice", Age: 30}
	fmt.Println(p)

	// Pointer type
	var ptr *int = &i
	fmt.Println(ptr)

	// Function type
	var fn func(int) int = func(x int) int { return x * x }
	fmt.Println(fn(3))

	// Interface type
	var any interface{} = "anything"
	fmt.Println(any)

	fmt.Printf("%T\n", ptr)
}
