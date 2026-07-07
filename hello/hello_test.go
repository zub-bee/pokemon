package greet

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("specify greeting in language requested", func(t *testing.T) {
		t.Run("spanish", func(t *testing.T) {
			got := Hello("Chris", "Spanish")
			want := "Hola, Chris"
			assertCorrectMessage(t, got, want)
		})
		t.Run("french", func(t *testing.T) {
			got := Hello("Chris", "French")
			want := "Bonjour, Chris"

			assertCorrectMessage(t, got, want)
		})
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
