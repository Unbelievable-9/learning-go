package main

import "fmt"

func main() {
	// Total & Average
	array := [5]int{
		1,
		2,
		3,
		4,
		5, // Needs comma on end, unusual
	}

	slice := array[1:] // low includes index but high doesn't

	totalAndAverage(slice)

	// Slice append
	appendSlice := append(slice, 6, 7)

	totalAndAverage(appendSlice)

	// Slice copy
	copySlice := make([]int, 2)
	copy(copySlice, appendSlice)

	totalAndAverage(copySlice)

	// Map inital
	goMap := make(map[string]string)
	goMap["name"] = "Jack"
	goMap["age"] = "27"
	goMap["gender"] = "male"

	fmt.Printf("The value for key 'name' is: %s\n", goMap["name"])

	// Map delete
	delete(goMap, "age")
	fmt.Printf("Length of the map is: %d\n", len(goMap))

	// Map key check
	if _, ok := goMap["test"]; ok {
		fmt.Printf("The value for key 'test' is: %s\n", goMap["test"])
	} else {
		fmt.Println("No key in map.")
	}

	// Better map
	bettermap := map[string]map[string]int{
		"A": map[string]int{
			"Alice": 18,
			"Aaron": 19,
		},
		"B": map[string]int{
			"Bob":   20,
			"Baker": 21,
		},
	}

	for key, subMap := range bettermap {
		for subKey, value := range subMap {
			fmt.Printf("Key: %s Subkey: %s Value: %d\n", key, subKey, value)
		}
	}

	if element, ok := bettermap["A"]; ok {
		if subElement, ok := element["Alice"]; ok {
			fmt.Printf("Alice is in category A and aged %d.\n", subElement)
		}
	}
	// Exercise 4
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	findTheSmallest(x)
}

// totalAndAverage calculates the total and average of an array
func totalAndAverage(slice []int) {
	total := 0
	average := 0.0

	// Underscore means 'ignore'
	for _, value := range slice {
		total += value
	}

	average = float64(float64(total) / float64(len(slice)))

	fmt.Println(slice)
	fmt.Printf("Total of slice is: %d\n", total)
	fmt.Printf("Average of slice is: %.2f\n", average)
}

func findTheSmallest(slice []int) {
	var smallest int

	if len(slice) <= 0 {
		smallest = -1
	} else {
		for index, value := range slice {
			if index == 0 || value < smallest {
				smallest = value
			}
		}
	}

	fmt.Printf("The smallest number in array is: %d\n", smallest)
}
