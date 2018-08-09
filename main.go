package main

import "fmt"

func main() {
	fmt.Print("Enter Temperture: ")
	var temp float32
	fmt.Scanf("%f", &temp)

	fmt.Printf("The Celsius Temperture of Fahrenheit %.1f is: %.1f", temp, fahrenheitToCelsius(temp))
}

// fahrenheitToCelsius is for test
func fahrenheitToCelsius(temp float32) float32 {
	return (temp - 32) * 5 / 9
}
