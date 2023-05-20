package peri

import "testing"

func TestPerimeter(t *testing.T) {

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
}

func TestArea(t *testing.T) {
	//table driven test
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
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

func errCheck(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
