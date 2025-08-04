package main


type Complex struct {
	real      float64
	imaginary float64
}

func newComplex(real float64, imag float64) *Complex {
	c := Complex{
		real:      real,
		imaginary: imag,
	}
	return &c
}
