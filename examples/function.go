package examples

import "fmt"

func addInStyle(args ...int) int {
	total := 0

	for _, value := range args {
		total += value
	}

	return total
}

func dinningPhilosopher() {
	panic("Dinning Philosopher is a problem you never solved.")
}

// FunctionExample runs examples about functions
func FunctionExample() {

	// Array
	fmt.Printf("The total of numbers under 3 is: %d\n", addInStyle(1, 2))

	// Slice
	intSlice := []int{1, 2, 3, 4}
	fmt.Printf("The total of numbers under 5 is: %d\n", addInStyle(intSlice...))
}
