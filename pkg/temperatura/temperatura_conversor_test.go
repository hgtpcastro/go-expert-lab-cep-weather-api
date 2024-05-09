package temperatura

import (
	"testing"
)

func TestConversor_CelsiusToFahrenheit(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "should convert 0 celsius to 32 fahrenheit",
			args: args{v: 0},
			want: 32,
		},
		{
			name: "should convert 100 celsius to 212 fahrenheit",
			args: args{v: 100},
			want: 212,
		},
		{
			name: "should convert -40 celsius to -40 fahrenheit",
			args: args{v: -40},
			want: -40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConversor()
			if got := c.CelsiusToFahrenheit(tt.args.v); got != tt.want {
				t.Errorf("Conversor.CelsiusToFahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConversor_CelsiusToKelvin(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "should convert 0 celsius to 273.15 kelvin",
			args: args{v: 0},
			want: 273.15,
		},
		{
			name: "should convert 100 celsius to 373.15 kelvin",
			args: args{v: 100},
			want: 373.15,
		},
		{
			name: "should convert -40 celsius to 233.15 kelvin",
			args: args{v: -40},
			want: 233.15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConversor()
			if got := c.CelsiusToKelvin(tt.args.v); got != tt.want {
				t.Errorf("Conversor.CelsiusToKelvin() = %v, want %v", got, tt.want)
			}
		})
	}
}
