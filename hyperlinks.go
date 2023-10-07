package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

// Link creates a terminal Hyperlink that can be clicked, if the terminal
// emulator supports it.
func Link(target url.URL, text string) string {
	return fmt.Sprintf(hyperlink, target.String(), text)
}

func LinkString(target string, text string) string {
	return fmt.Sprintf(hyperlink, target, text)
}

func LinkFile(file string, label string) string {
	hostname, _ := os.Hostname() // no error checking as most clients also accept file:// links without hostname

	path, err := filepath.Abs(file)
	if err != nil {
		path = file
	}

	u := url.URL{
		Scheme: "file",
		Host:   hostname,
		Path:   path,
	}
	return Link(u, label)
}
