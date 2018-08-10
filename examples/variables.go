package main

import (
	"fmt"
	"math"
)

var (
	a = "This"
	b = "is"
	c = "how"
	d = "you"
	e = "define"
	f = "multiple"
	g = "variables."
)

const (
	h = "Const"
	i = "is"
	j = "the"
	k = "same.\n"
)

func main() {
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(h, i, j, k)

	fmt.Print("Enter Temperture: ")
	var temp float64
	fmt.Scanf("%f", &temp)

	fmt.Printf("The Celsius Temperture of Fahrenheit %.1f is: %.1f\n\n", temp, fahrenheitToCelsius(temp))

	fmt.Print("Enter Feet:")
	var feet float64
	fmt.Scanf("%f", &feet)

	fmt.Printf("%.1f feet(s) is equal to %.2f meter(s)", feet, feetToMeter(feet))
}

// fahrenheitToCelsius is for test
func fahrenheitToCelsius(temp float64) float64 {
	temp = math.Ceil(temp*100) / 100

	return (temp - 32) * 5 / 9
}

// feetToMeter is for test
func feetToMeter(feet float64) float64 {
	feet = math.Ceil(feet*100) / 100

	return feet * 0.3048
}
