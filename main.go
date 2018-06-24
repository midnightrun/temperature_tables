package main

import (
	"fmt"
	"strings"
	"strconv"
)

//type Celsius float64
//type Kelvin float64

type conversion func(v float64) float64

// CelsiusToKelvin converts C To K
func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

// kelvinTo converts C To K
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func drawTable(c conversion) {
	seperator := strings.Repeat("=", 27)
	fmt.Println(seperator)
	fmt.Printf("| %-10v | %-10v |\n", "C", "K")
	fmt.Println(seperator)

	for i:=-40.0; i <= 100; i+=5 {
		v := strconv.FormatFloat(c(i), 'f', 2, 64)
		fmt.Printf("| %-10v | %-10v |\n", i, v)
	}

	fmt.Println(seperator)
}

func main() {
	drawTable(celsiusToKelvin)
}