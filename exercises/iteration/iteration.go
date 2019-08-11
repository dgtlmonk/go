package interation

const repeatCount = 5

// Repeat repeats a string
func Repeat(char string, howMany int) string {

	var repeated string

	for i := 0; i < howMany; i++ {
		repeated += char
	}

	return repeated
}
