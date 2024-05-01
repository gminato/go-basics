package iteration

import "testing"

func TestFor(t *testing.T) {
	t.Run("testing for repetation", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q and got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("B")
	}
}
