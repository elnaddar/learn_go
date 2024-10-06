package main

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := rec.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rec := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, rec, want)
	})

	t.Run("circle", func(t *testing.T) {
		crcl := Circle{10}
		want := 314.1592653589793
		checkArea(t, crcl, want)
	})
}
