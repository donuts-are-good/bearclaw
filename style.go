package main

import "fmt"

// Bold is equivalent to SetBold() + content + UnsetBold(), content will
// thereby be printed in bold.
func Bold(content ...interface{}) string {
	return SetBold() + fmt.Sprint(content...) + UnsetBold()
}

// SetBold makes the following text become bold
func SetBold() string {
	return ansiescape + bold + set
}

// UnsetBold removes the bold property
func UnsetBold() string {
	return ansiescape + boldOff + set
}

// Faint is equivalent to SetFaint() + content + UnsetFaint(), content will
// thereby be printed in a less intense color.
func Faint(content ...interface{}) string {
	return SetFaint() + fmt.Sprint(content...) + UnsetFaint()
}

// SetFaint makes the color of the following text less bright
func SetFaint() string {
	return ansiescape + faint + set
}

// UnsetFaint removes the faint property
func UnsetFaint() string {
	return ansiescape + boldOff + set
}

// Italic is equivalent to SetItalic() + content + UnsetItalic(), content will
// thereby be italic.
func Italic(content ...interface{}) string {
	return SetItalic() + fmt.Sprint(content...) + UnsetItalic()
}

// SetItalic makes the text following text become italic
func SetItalic() string {
	return ansiescape + italic + set
}

// UnsetItalic removes the italic property
func UnsetItalic() string {
	return ansiescape + italicOff + set
}

// Underscore is equivalent to SetUnderscore() + content + UnsetUnderscore(),
// content will thereby be underscored.
func Underscore(content ...interface{}) string {
	return SetUnderscore() + fmt.Sprint(content...) + UnsetUnderscore()
}

// SetUnderscore makes the text following become underlined
func SetUnderscore() string {
	return ansiescape + underscore + set
}

// UnsetUnderscore removes the underscore property
func UnsetUnderscore() string {
	return ansiescape + underscoreOff + set
}

// DoubleUnderscore is equivalent to SetDoubleUnderscore() + content +
// UnsetDoubleUnderscore(), content will thereby be double-underscored.
func DoubleUnderscore(content ...interface{}) string {
	return SetDoubleUnderscore() + fmt.Sprint(content...) + UnsetDoubleUnderscore()
}

// SetDoubleUnderscore makes the following text become underscored
func SetDoubleUnderscore() string {
	return ansiescape + doubleUnderscore + set
}

// UnsetDoubleUnderscore removes the double-underscore attribute
func UnsetDoubleUnderscore() string {
	return ansiescape + underscoreOff + set
}

// Blink is equivalent to SetBlink() + content + UnsetBlink(), content will
// thereby blink.
func Blink(content ...interface{}) string {
	return SetBlink() + fmt.Sprint(content...) + UnsetBlink()
}

// SetBlink makes the following text blink
func SetBlink() string {
	return ansiescape + blink + set
}

// UnsetBlink stops the text from blinking
func UnsetBlink() string {
	return ansiescape + blinkOff + set
}

// FastBlink is equivalent to SetFastBlink() + content + UnsetFastBlink(),
// content will thereby blink fast.
func FastBlink(content ...interface{}) string {
	return SetFastBlink() + fmt.Sprint(content...) + UnsetFastBlink()
}

// SetFastBlink makes the following text blink fast
func SetFastBlink() string {
	return ansiescape + fastblink + set
}

// UnsetFastBlink stops the text from blinking fast
func UnsetFastBlink() string {
	return ansiescape + fastblinkOff + set
}

// ReverseVideo is equivalent to SetReverseVideo() + content +
// UnsetReverseVideo(), content will thereby be inverted. This means that the
// background color and the foreground color will be switched.
func ReverseVideo(content ...interface{}) string {
	return SetReverseVideo() + fmt.Sprint(content...) + UnsetReverseVideo()
}

// SetReverseVideo makes the following text become inverted.
func SetReverseVideo() string {
	return ansiescape + reverseVideo + set
}

// UnsetReverseVideo disables the inversion
func UnsetReverseVideo() string {
	return ansiescape + reverseVideoOff + set
}

// Conceal is equivalent to SetConceal() + content + UnsetConceal(), content
// will thereby be displayed as spaces. It will still be visible if written to
// a file
func Conceal(content ...interface{}) string {
	return SetConceal() + fmt.Sprint(content...) + UnsetConceal()
}

// SetConceal conceals the following text
func SetConceal() string {
	return ansiescape + conceal + set
}

// UnsetConceal removes the concealed property
func UnsetConceal() string {
	return ansiescape + concealOff + set
}

// Strikethrough is equivalent to SetStrikethrough() + content +
// UnsetStrikethrough(), content will thereby be displayed as crossed out
func Strikethrough(content ...interface{}) string {
	return SetStrikethrough() + fmt.Sprint(content...) + UnsetStrikethrough()
}

// SetStrikethrough makes the following text crossed out
func SetStrikethrough() string {
	return ansiescape + strikethrough + set
}

// UnsetStrikethrough removes the strikethrough property
func UnsetStrikethrough() string {
	return ansiescape + strikethroughOff + set
}

// Reset clears all colors and styles (bold, underscore, blink, concealed)
func Reset() string {
	return ansiescape + reset + set
}
