package calculator

import (
	"errors"
	"math"
)


func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of a negative not allowed")
	}
	var z float64 = 1
	for i:=1;i<=10; i++{
		z = (z-(math.Pow(z,2)-a)/(2*z))
	}
	return z, nil
}