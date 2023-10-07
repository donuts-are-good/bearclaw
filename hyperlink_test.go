package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestLink(t *testing.T) {
	t.Run("url", func(t *testing.T) {
		u, _ := url.Parse("https://example.com")

		fmt.Printf("Link to example.com: %s\n", Link(*u, "Linktext"))
	})
	t.Run("string", func(t *testing.T) {
		fmt.Printf("Link to my homepage: %s\n", LinkString("https://moritz.sh", "Linktext"))
	})
	t.Run("file", func(t *testing.T) {
		fmt.Printf("Link to hyperlink_test.go: %s\n", LinkFile("hyperlink_test.go", "Linktext"))
	})
}
