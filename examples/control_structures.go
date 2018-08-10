package main

import "fmt"

func main() {
	// Times Table
	printTimesTable()

	fmt.Println()

	// Have fun with 'switch'
	fmt.Println("Input a number within 0-9")
	var number int
	fmt.Scanf("%d", &number)

	word, status := numberToWord(number)

	if status {
		fmt.Printf("Number %d in English is: %s\n", number, word)
	} else {
		fmt.Println(word)
	}

	fmt.Println()

	// Exercise 2
	divisibelByThree()

	fmt.Println()

	// Exercise 3
	fizzBuzz()
}

// printTimesTable just print a simple Times Table
func printTimesTable() {
	fmt.Println("Times Table:")
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			if (i > 2 && i < 5) && j == 2 {
				fmt.Printf("%d * %d = %d  ", j, i, j*i)
			} else {
				fmt.Printf("%d * %d = %d ", j, i, j*i)
			}
		}

		fmt.Println()
	}
}

// numberToWord is a dummy function only accept number within 0-9
func numberToWord(number int) (string, bool) {
	var word string
	status := true

	switch number {
	case 1:
		word = "One"
	case 2:
		word = "Two"
	case 3:
		word = "Three"
	case 4:
		word = "Four"
	case 5:
		word = "Five"
	case 6:
		word = "Six"
	case 7:
		word = "Seven"
	case 8:
		word = "Eight"
	case 9:
		word = "Nine"
	case 0:
		word = "Zero"
	default:
		word = "Wrong input number."
		status = false
	}

	return word, status
}

// divisibelByThree print all numbers in 1-100 divisable by 3
func divisibelByThree() {
	fmt.Println("Numbers divisible by 3:")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println()
}

// fizzBuzz print all numbers in 1-100
// multiples of three are relpaced by "Fizz"
// multiples of five are relpaced by "Buzz"
// multiples of three and five are replaced bu "FizzBuzz"
func fizzBuzz() {
	fmt.Println("Fizz Buzz:")
	for i := 1; i <= 100; i++ {
		multipleOfThree := false
		multipleOfFive := false

		if i%3 == 0 {
			multipleOfThree = true
		}

		if i%5 == 0 {
			multipleOfFive = true
		}

		if multipleOfThree && multipleOfFive {
			fmt.Print("FizzBuzz ")
		} else if multipleOfThree {
			fmt.Print("Fizz ")
		} else if multipleOfFive {
			fmt.Print("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}
