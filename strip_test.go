package main

import (
	"fmt"
	"regexp"
	"testing"
)

func BenchmarkStripping(b *testing.B) {
	DetectionPattern := regexp.MustCompile(`(?m)(((\x1b\[[;\d]*[A-Za-z])|\x1b]8;;|\x1b\\.*?\x1b|]8;;\x1b\\)*)`)
	str := fmt.Sprint(Black("black text"), "some kept text", ReverseVideo("inverted text"), Green("green text"), Bold("bold text"))

	b.Run("regex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DetectionPattern.ReplaceAllString(str, "")
		}
	})

	b.Run("split", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StripString(str)
		}
	})
}

func TestStrip(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "formatting",
			input:    fmt.Sprint(Black("black text"), "some kept text", ReverseVideo("inverted text"), Green("green text"), Bold("bold text")),
			expected: "black textsome kept textinverted textgreen textbold text",
		},
		{
			name:     "link",
			input:    fmt.Sprintf("Link to my homepage: %s", LinkString("https://moritz.sh", "Linktext")),
			expected: "Link to my homepage: https://moritz.sh",
		},
		{
			name:     "notification",
			input:    fmt.Sprintf("Send a notification %sto me", SendNotification("title", "body")),
			expected: "Send a notification to me",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if StripString(test.input) != test.expected {
				t.Logf("got:      %s\n", StripString(test.input))
				t.Logf("expected: %s\n", test.expected)
				t.Log([]byte(StripString(test.input)))
				t.Log([]byte(test.expected))
			}
		})
	}
}
