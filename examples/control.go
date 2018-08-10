package main

import "fmt"

func main() {
	printTimesTable()
}

func printTimesTable() {
	j := 1

	for i := 1; i < 10; i++ {
		for j < i+1 {
			if (i > 2 && i < 5) && j == 2 {
				fmt.Printf("%d * %d = %d  ", j, i, j*i)
			} else {
				fmt.Printf("%d * %d = %d ", j, i, j*i)
			}

			j++
		}

		j = 1
		fmt.Println()
	}
}
