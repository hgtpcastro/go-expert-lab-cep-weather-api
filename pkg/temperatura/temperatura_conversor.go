package temperatura

type conversor struct {
}

func NewConversor() *conversor {
	return &conversor{}
}

func (c *conversor) CelsiusToFahrenheit(v float64) float64 {
	return v*1.8 + 32
}

func (c *conversor) CelsiusToKelvin(v float64) float64 {
	return v + 273.15
}
