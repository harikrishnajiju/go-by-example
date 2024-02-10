package main

import (
	"flag"
	"fmt"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	// Define flags
	var (
		celsiusFlag    = flag.Float64("celsius", 0.0, "Temperature in Celsius")
		fahrenheitFlag = flag.Float64("fahrenheit", 0.0, "Temperature in Fahrenheit")
	)

	// Parse command line arguments
	flag.Parse()

	// Check if both flags are provided
	if *celsiusFlag != 0 && *fahrenheitFlag != 0 {
		fmt.Println("Please provide only one temperature value.")
		return
	}

	// Perform conversion based on the provided flag
	if *celsiusFlag != 0 {
		fahrenheit := celsiusToFahrenheit(*celsiusFlag)
		fmt.Printf("%.2f°C is equal to %.2f°F\n", *celsiusFlag, fahrenheit)
	} else if *fahrenheitFlag != 0 {
		celsius := fahrenheitToCelsius(*fahrenheitFlag)
		fmt.Printf("%.2f°F is equal to %.2f°C\n", *fahrenheitFlag, celsius)
	} else {
		fmt.Println("Please provide a temperature value using either -celsius or -fahrenheit flag.")
	}
}
