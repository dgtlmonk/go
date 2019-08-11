package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		num := []int{1, 10, 4}
		got := Sum(num)
		want := 15

		if got != want {
			t.Errorf("got %q want %q given, %v", got, want, num)
		}
	})
}

// Note: use $ test -cover
