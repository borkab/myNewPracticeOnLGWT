package peri

import "testing"

func TestPerimeter(t *testing.T) {
	//TDT
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g tt.want %g", got, tt.want)
		}
	}

	/*
		checkPerimeter := func(t testing.TB, sh Shape, want float64) {
			t.Helper()
			got := sh.Perimeter()
			errCheck(t, got, want)
		}
		t.Run("rectangle", func(t *testing.T) {
			rec := Rectangle{10.0, 10.0}
			checkPerimeter(t, rec, 40.0)
		})

		t.Run("circle", func(t *testing.T) {
			cir := Circle{10}
			checkPerimeter(t, cir, 62.83185307179586)
		})

		t.Run("triangle", func(t *testing.T) {
			tri := Triangle{12, 6}
			checkPerimeter(t, tri, 36.0)
		})
	*/
}

func TestArea(t *testing.T) {
	//table driven test
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Baseline: 12, Height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
	/*
		checkArea := func(t testing.TB, sh Shape, want float64) {
			t.Helper()
			got := sh.Area()
			errCheck(t, got, want)
		}

		t.Run("rectangle", func(t *testing.T) {
			rec := Rectangle{12.0, 6.0}
			checkArea(t, rec, 72.0)
		})

		t.Run("circle", func(t *testing.T) {
			cir := Circle{10}
			checkArea(t, cir, 314.1592653589793)
			//pi: 3.141592653589793

		t.Run("triangle", func(t *testing.T) {
			tri:= Triangle{12, 6}
			checkArea(t, tri, 36.0)
		})

		})
	*/

}

/*
func errCheck(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
*/
