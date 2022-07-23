package conversion

type UnitConversion interface {
	ConvertToCelsius(arg float64) float64
	ConvertToFahrenheit(arg float64) float64
	ConvertToKelvin(arg float64) float64
}

type Celsius struct {
	arg float64
}

func (d Celsius) ConvertToFahrenheit(v float64) float64 {
	return v*1.8 + 32
}

func (d Celsius) ConvertToKelvin(v float64) float64 {
	return v + 273.5
}

type Fahrenheit struct {
	arg float64
}

func (d Fahrenheit) ConvertToKelvin(v float64) float64 {
	return (459.67 + v) * 5 / 9
}

func (d Fahrenheit) ConvertToCelsius(v float64) float64 {
	return (v - 32) * .556
}

type Kelvin struct {
	arg float64
}

func (d Kelvin) ConvertToFahrenheit(v float64) float64 {
	return (v-273.15)*9/5 + 32
}

func (d Kelvin) ConvertToCelsius(v float64) float64 {
	return v - 273.15
}
