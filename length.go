package main

import (
	"regexp"
	"unicode/utf8"
)

// GetLengthWithoutCodes counts the characters after all Escape-sequences have
// been removed.
func GetLengthWithoutCodes(content string) int {
	re := regexp.MustCompile(`(?m)\x1b(\[([\d;]+m|[\d;]+(H|A|B|C|D|J|K|S|T)|s|u)|\](8;;[^\x1b]*\x1b\\|777;notify;[^;]*;[^\a]*\a))`)
	return utf8.RuneCountInString(re.ReplaceAllString(content, ""))
}
