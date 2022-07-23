package main

import (
	"fmt"

	conversion "./conversion/"
)

func main() {
	fmt.Println("")
	var fConversion conversion.Fahrenheit

	fConversion.ConvertToCelsius(25)
	fConversion.ConvertToKelvin(25)

	var kConversion conversion.Kelvin
	kConversion.ConvertToCelsius(25)
	kConversion.ConvertToFahrenheit(25)

	var cConversion conversion.Celsius

	cConversion.ConvertToFahrenheit(25)
	cConversion.ConvertToKelvin(25)
}
