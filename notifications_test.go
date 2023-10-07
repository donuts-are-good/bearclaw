package main

import (
	"fmt"
	"testing"
)

func TestNotification(t *testing.T) {
	fmt.Printf("Here should be no text, if your terminal supports notifications: %s\n", SendNotification("Hello from go-ansi", "This notification should be sent through your terminal emulator, it it's supported."))
}
