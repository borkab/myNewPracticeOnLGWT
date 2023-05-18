package peri

type Rectangle struct {
	Height float64
	Width  float64
}

func Perimeter(height, width float64) float64 {
	return 2 * (height + width)
}

func Area(height, width float64) float64 {
	return height * width
}
