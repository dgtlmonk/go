package person

import "testing"

func TestPerson(t *testing.T) {
	assertGender := func(t *testing.T, m Mamal, gender string) {
		t.Helper()
		got := m.Gender()

		if got != gender {
			t.Errorf("got %s want %s", got, gender)
		}
	}

	t.Run("person should have correct gender", func(t *testing.T) {
		person := Person{"Male"}
		want := "Male"

		assertGender(t, &person, want)
	})

	t.Run("animal should have correct gender", func(t *testing.T) {
		animal := Animal{"Female"}
		want := "Female"

		assertGender(t, animal, want)
	})
}
