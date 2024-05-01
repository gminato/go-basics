package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("adding two integer", func(t *testing.T) {
		got := Add(2, 3)
		want := 5
		if got != want {
			t.Errorf("wanted %d got %d", want, got)
		}
	})
}
