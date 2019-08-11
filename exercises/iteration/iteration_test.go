package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertRepeatTimes := func(t *testing.T, repeated, expected string) {
		/*
			t.Helper() is needed to tell the test suite that this method is a helper.
			By doing this when it fails the line number reported
			will be in our function call rather than inside our test helper.
		*/
		t.Helper()

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeating 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		assertRepeatTimes(t, repeated, expected)
	})

	t.Run("repeating 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertRepeatTimes(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeat := Repeat("b", 3)
	fmt.Println(repeat)
	// Output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}
