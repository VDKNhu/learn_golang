package more_types

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	// pointer
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer -> 42
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i -> 21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j -> 73

	// pointer to struct
	v := Vertex{1, 2}
	p2 := &v
	p2.X = 1e9
	fmt.Println(v) // {1000000000 2}

	// array + slice
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]

	// slice length and capacity
	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)

	// Slice the slice to give it zero length.
	s2 = s2[:0]
	printSlice(s2)

	// Extend its length.
	s2 = s2[:4]
	printSlice(s2)

	// Drop its first two values.
	s2 = s2[2:]
	printSlice(s2)

	// exercise slices
	pic.Show(Pic)

	// exercise maps
	wc.Test(WordCount)

	// exercise closure
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy, dy)

	for y := range pic {
		pic[y] = make([]uint8, dx)

		for x := range pic[y] {
			pic[y][x] = uint8(x * x * y * y)
		}
	}

	return pic
}

func WordCount(s string) map[string]int {
	strParts := strings.Split(s, " ")
	result := make(map[string]int)

	for _, value := range strParts {
		_, isIn := result[value]
		if !isIn {
			result[value] = 0
		}

		result[value]++
	}

	return result
}

func fibonacci() func(int) int {
	fibResult := 0

	return func(value int) int {
		fibResult += value
		return fibResult
	}
}

// pointer
// A pointer holds the memory address of a value.
// The & operator generates a pointer to its operand.
// The * operator denotes the pointer's underlying value.

// array
// The type [n]T is an array of n values of type T.
// Arrays cannot be resized

// slice
// The type []T is a slice with elements of type T.
// A slice, is a dynamically-sized, flexible view into the elements of an array.
// Slices are like REFERENCES to arrays. Changing the elements of a slice modifies the corresponding elements of its underlying array.
// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
// a := make([]int, 5) // len(a)=5
// b := make([]int, 0, 5) // len(b)=0, cap(b)=5

// map
// A map maps keys to values.
// The type make(map[key_type]value_type) is a map with key type is key_type and value type is value_type.

// closure
// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
