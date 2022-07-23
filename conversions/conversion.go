package conversion

type UnitConversion interface {
	ConvertToCelsius(arg float64) float64
	ConvertToFahrenheit(arg float64) float64
	ConvertToKelvin(arg float64) float64
}

type Celsius struct {
	Name string
}

func (v Celsius) ConvertToFahrenheit(arg float64) float64 {
	return fahrenheitConversion(arg)
}

func (v Celsius) ConvertToKelvin(arg float64) float64 {
	return kelvinConversion(arg)
}

type Fahrenheit struct {
	Name string
}

func (v Fahrenheit) ConvertToFahrenheit(arg float64) float64 {
	return fahrenheitConversion(arg)
}

func (v Fahrenheit) ConvertToCelsius(arg float64) float64 {
	return celsiusConversion(arg)
}

type Kelvin struct {
	Name string
}

func (v Kelvin) ConvertToFahrenheit(arg float64) float64 {
	return fahrenheitConversion(arg)
}

func (v Kelvin) ConvertToCelsius(arg float64) float64 {
	return celsiusConversion(arg)
}

func kelvinConversion(arg float64) float64 {
	return 0.0

}
func fahrenheitConversion(arg float64) float64 {
	return 0.0

}
func celsiusConversion(arg float64) float64 {
	return 0.0

}
