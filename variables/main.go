package main

import "fmt"

func main() {
	// 1. Declare a single variable with a specific type
	var a int
	a = 10

	// 2. Declare and initialize a single variable with a specific type
	var b int = 20

	// 3. Declare and initialize a single variable with type inference
	var c = 30

	// 4. Short variable declaration (only inside functions)
	d := 40

	// 5. Declare multiple variables with specific types
	var e, f int
	e = 50
	f = 60

	// 6. Declare and initialize multiple variables with specific types
	var g, h int = 70, 80

	// 7. Declare and initialize multiple variables with type inference
	var i, j = 90, 100

	// 8. Short variable declaration for multiple variables (only inside functions)
	k, l := 110, 120

	// 9. Declare a variable with zero value
	var m string

	// 10. Declare and initialize a variable with a specific type and value
	var n string = "hello"

	// 11. Declare and initialize a variable with type inference
	var o = "world"

	// 12. Short variable declaration for a string (only inside functions)
	p := "Go"

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
}
