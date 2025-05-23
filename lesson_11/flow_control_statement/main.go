package flow_control_statement

import (
	"fmt"
)

func main() {
	// example 1:
	fmt.Println(Sqrt(2))

	// example 2:
	defer fmt.Println("world")
	fmt.Println("hello")

	// example 3:
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func Sqrt(x float64) float64 {
	prev, z := float64(0), float64(1)
	for abs(prev-z) > 1e-8 {
		prev = z
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func abs(number float64) float64 {
	if number < 0 {
		return number * -1
	} else {
		return number
	}
}

// defer function
// A defer statement defers the execution of a function until the surrounding function returns
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
// example 2: hello world
// example 3: counting done 9 8 7 6 5 4 3 2 1 0
