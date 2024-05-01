package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		want := 15
		got := Sum(numbers)

		if want != got {
			t.Errorf("wanted %d got %d numbers %v ", want, got, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 5}

		want := 8
		got := Sum(numbers)

		if want != got {
			t.Errorf("wanted %d got %d numbers %v ", want, got, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
