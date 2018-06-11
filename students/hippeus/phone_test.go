package phone

import (
	"testing"
)

var tests = []struct {
	input, expected string
}{
	{"1234567890", "1234567890"},
	{"123 456 7891", "1234567891"},
	{"(123) 456 7892", "1234567892"},
	{"(123) 456-7893", "1234567893"},
	{"123-456-7894", "1234567894"},
	{"123-456-7890", "1234567890"},
	{"1234567892", "1234567892"},
	{"(123)456-7892", "1234567892"},
}

func TestNormalize(t *testing.T) {
	var got string
	for _, tt := range tests {
		got = Normalize(tt.input)
		if got != tt.expected {
			t.Fatalf("got: %s; want: %s\n", got, tt.expected)
		}
	}
}
