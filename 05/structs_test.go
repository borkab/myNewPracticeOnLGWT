package peri

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	want := 40.0

	errCheck(t, got, want)
}

func TestArea(t *testing.T) {
	rec := Rectangle{}
	got := Area(12.0, 6.0)
	want := 72.0

	errCheck(t, got, want)
}

func errCheck(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}