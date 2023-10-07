package main

import (
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedLength int
	}{
		{
			name:           "color",
			input:          Blue("4444"),
			expectedLength: 4,
		},
		{
			name:           "formatting",
			input:          Bold("88888888"),
			expectedLength: 8,
		},
		{
			name:           "link",
			input:          LinkString("https://example.com", "333"),
			expectedLength: 3,
		},
		{
			name:           "notification",
			input:          SendNotification("some title", "some body"),
			expectedLength: 0,
		},
		{
			name:           "movement",
			input:          fmt.Sprintf("Up%s3x%s", Up(), UpX(400)),
			expectedLength: 4,
		},
		{
			name:           "cursor-save",
			input:          fmt.Sprintf("Save%sRestore%s:)", SavePos(), RestorePos()),
			expectedLength: 13,
		},
		{
			name:           "clear-screen",
			input:          fmt.Sprintf("clear%s", ClearScreen()),
			expectedLength: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if l := GetLengthWithoutCodes(test.input); l != test.expectedLength {
				t.Logf("expected length %d, but got %d", test.expectedLength, l)
				fmt.Printf("%x", GetLengthWithoutCodes(test.input))
				t.Fail()
			}
		})
	}
}
