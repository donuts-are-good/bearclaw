package main

import (
	"fmt"
)

// SendNotification instructs the terminal emulator to send a notification to
// the user. This will also work via SSH connections. Support for this code is
// not ubiquitous and may lead to messed up output.
//
// The provided title may not contain semi-colon. Everything after the first
// semicolon will be part of the body.
func SendNotification(title, body string) string {
	return fmt.Sprintf(notification, title, body)
}
