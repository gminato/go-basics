package stucts

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}

	}

	t.Run("rectangle", func(t *testing.T) {
		rectange := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, &rectange, want)

	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, &circle, want)
	})

}
