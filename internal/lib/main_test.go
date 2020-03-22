package lib_test

import (
	"testing"

	"github.com/stang/ci-sandbox/internal/lib"
)

func TestGreetings(t *testing.T) {
	tests := []struct {
		name, lang, expected string
		err                  error
	}{
		{"World", "en", "Hello World!", nil},
		{"john", "en", "Hello John!", nil},
		{"", "en", "Hello Everyone!", nil},
		{"john", "fr", "Bonjour John!", nil},
		{"", "fr", "Bonjour Tout Le Monde!", nil},
	}

	for _, test := range tests {
		res, err := lib.Greetings(test.lang, test.name)
		if err != test.err {
			t.Errorf("Greetings returned unexpected error for '%+v': '%s'", test, err)
		}
		if res != test.expected {
			t.Errorf("Greetings returned unexpected value for '(%s, %s)': got '%s' instead of '%s'", test.name, test.lang, res, test.expected)
		}
	}
}
