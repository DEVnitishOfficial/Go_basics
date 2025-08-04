package app

import (
	"fmt"
	"strconv"
)

type Complex struct {
	Real      float64
	Imaginary float64
}

func NewComplex(real float64, imag float64) *Complex {
	ComplexNumber := Complex{
		Real:      real,
		Imaginary: imag,
	}
	return &ComplexNumber
}

// Add two complex numbers
func (c1 Complex) Add(c2 Complex) Complex {
	return Complex{
		Real:      c1.Real + c2.Real,
		Imaginary: c1.Imaginary + c2.Imaginary,
	}
}

// Subtract two complex numbers
func (c1 Complex) Subtract(c2 Complex) Complex {
	return Complex{
		Real:      c1.Real - c2.Real,
		Imaginary: c1.Imaginary - c2.Imaginary,
	}
}

// multiply two complex numbers
func (c1 Complex) Multiply(c2 Complex) Complex {
	return Complex{
		Real:      c1.Real*c2.Real - c1.Imaginary*c2.Imaginary,
		Imaginary: c1.Real*c2.Imaginary + c1.Imaginary*c2.Real,
	}
}

// Divide two complex numbers
func (c1 Complex) Divide(c2 Complex) (Complex, error) {
	if c2.Real == 0 && c2.Imaginary == 0 {
		return Complex{}, fmt.Errorf("cannot divide by zero")
	}
	denominator := c2.Real*c2.Real + c2.Imaginary*c2.Imaginary
	return Complex{
		Real:      (c1.Real*c2.Real + c1.Imaginary*c2.Imaginary) / denominator,
		Imaginary: (c1.Imaginary*c2.Real - c1.Real*c2.Imaginary) / denominator,
	}, nil
}

// Display method to show the complex number in readable format
func (c Complex) Display() string {
	realStr := strconv.FormatFloat(c.Real, 'f', 2, 64)
	imagStr := strconv.FormatFloat(c.Imaginary, 'f', 2, 64)

	sign := "+"
	if c.Imaginary < 0 {
		sign = "-"
		imagStr = strconv.FormatFloat(-c.Imaginary, 'f', 2, 64)
	}

	return fmt.Sprintf("%s %s %si", realStr, sign, imagStr)
}
