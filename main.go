package main

import (
	"fmt"
	"strings"
)

type conversion func(v float64) float64

func fahrenheitToCelsius(v float64) float64 {
	return (v - 32.0) * 5.0 / 9.0
}

func celsiusToFahrenheit(v float64) float64 {
	return (v * 9 / 5) + 32
}

func main() {
	drawTable("C", "F", celsiusToFahrenheit)
	drawTable("F", "C", fahrenheitToCelsius)
}

func drawTable(source, destination string, c conversion) {
	// Draw header
	line := strings.Repeat("=", 41)
	fmt.Println(line)
	fmt.Printf("|°%-18v|°%-18v|\n", source, destination)
	fmt.Println(line)

	// Render Input
	for i := -40; i <= 100; i += 5 {
		fmt.Printf("|%-19v|%-19.4v|\n", i, c(float64(i)))
	}

	// Insert footer
	fmt.Println(line)
}
