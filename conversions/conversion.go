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
	//return math.Sin(arg)
}

func (v Celsius) ConvertToKelvin(arg float64) float64 {
	//	return math.Sin(arg)
}

type Fahrenheit struct {
	Name string
}

func (v Fahrenheit) ConvertToFahrenheit(arg float64) float64 {
	//return math.Sin(arg)
}

func (v Fahrenheit) ConvertToCelsius(arg float64) float64 {
	//return math.Sin(arg)
}

type Kelvin struct {
	Name string
}

func (v Kelvin) ConvertToFahrenheit(arg float64) float64 {
	//	return math.Sin(arg)
}

func (v Kelvin) ConvertToKelvin(arg float64) float64 {
	//	return math.Sin(arg)
}
