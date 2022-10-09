package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("Expected %d but got %d given %v", want, got, numbers)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum the tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %d but got %d ", want, got)
		}
	})

	t.Run("sum an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if!reflect.DeepEqual(got, want) {
			t.Errorf("Expected %d but got %d ", want, got)
		}
	})
}