package lib_test

import (
	"testing"

	"github.com/stang/ci-sandbox/internal/lib"
)

func TestGreetings(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"World", "Hello World!"},
		{"john", "Hello John!"},
		{"", "Hello Everyone!"},
	}

	for _, test := range tests {
		if res := lib.Greetings(test.input); res != test.expected {
			t.Errorf("Greetings returned unexpected value for '%s': got '%s' instead of '%s'", test.input, res, test.expected)
		}
	}
}
