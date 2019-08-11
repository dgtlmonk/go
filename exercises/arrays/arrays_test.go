package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		num := []int{1, 10, 4}
		got := Sum(num)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, num)
		}
	})

	t.Run("sum of two slices", func(t *testing.T) {
		slice1 := []int{2, 10}
		slice2 := []int{2, 10}

		got := SumAll(slice1, slice2)
		want := []int{12, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sum of all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

// Note: use $ test -cover
