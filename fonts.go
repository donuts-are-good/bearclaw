package main

import (
	"fmt"
	"strconv"
)

// Fraktur is equivalent to SetFraktur() + content + UnsetFraktur(), content
// will thereby be printed in Frakturschrift. This is very rarely support.
func Fraktur(content ...interface{}) string {
	return SetFraktur() + fmt.Sprint(content...) + UnsetFraktur()
}

// SetFraktur makes the following text be printed printed in Frakturschrift
func SetFraktur() string {
	return ansiescape + fraktur + set
}

// UnsetFraktur resets the font to the default
func UnsetFraktur() string {
	return ansiescape + frakturOff + set
}

// Font applies the specified fontface to the text
func Font(font int, content ...interface{}) string {
	return SetFont(font) + fmt.Sprint(content...) + UnsetFont()
}

// SetFont starts a block written in the specified font
func SetFont(font int) string {
	return ansiescape + strconv.Itoa(font+10) + set
}

// UnsetFont resets the fontface back to the default
func UnsetFont() string {
	return ansiescape + frakturOff + set
}
