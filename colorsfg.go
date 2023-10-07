package main

import "fmt"

// Black wraps the content in ANSI codes to make its foreground color black
func Black(content ...interface{}) string {
	return SetBlack() + fmt.Sprint(content...) + UnsetBlack()
}

// SetBlack sets the foreground color to black
func SetBlack() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + blackfg + set
}

// UnsetBlack resets the foreground color from black to default.
func UnsetBlack() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + resetfg + set
}

// Red wraps the content in ANSI codes to make its foreground color red
func Red(content ...interface{}) string {
	return SetRed() + fmt.Sprint(content...) + UnsetRed()
}

// SetRed sets the foreground color to red
func SetRed() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + redfg + set
}

// UnsetRed resets the foreground color from red to default.
var UnsetRed func() string = UnsetBlack

// Green wraps the content in ANSI codes to make its foreground color green
func Green(content ...interface{}) string {
	return SetGreen() + fmt.Sprint(content...) + UnsetGreen()
}

// SetGreen sets the foreground color to green
func SetGreen() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + greenfg + set
}

// UnsetGreen resets the foreground color from green to default.
var UnsetGreen func() string = UnsetBlack

// Yellow wraps the content in ANSI codes to make its foreground color yellow
func Yellow(content ...interface{}) string {
	return SetYellow() + fmt.Sprint(content...) + UnsetYellow()
}

// SetYellow sets the foreground color to yellow
func SetYellow() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + yellowfg + set
}

// UnsetYellow resets the foreground color from yellow to default.
var UnsetYellow func() string = UnsetBlack

// Blue wraps the content in ANSI codes to make its foreground color blue
func Blue(content ...interface{}) string {
	return SetBlue() + fmt.Sprint(content...) + UnsetBlue()
}

// SetBlue sets the foreground color to blue
func SetBlue() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + bluefg + set
}

// UnsetBlue resets the foreground color from blue to default.
var UnsetBlue func() string = UnsetBlack

// Magenta wraps the content in ANSI codes to make its foreground color magenta
func Magenta(content ...interface{}) string {
	return SetMagenta() + fmt.Sprint(content...) + UnsetMagenta()
}

// SetMagenta sets the foreground color to magenta
func SetMagenta() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + magentafg + set
}

// UnsetMagenta resets the foreground color from magenta to default.
var UnsetMagenta func() string = UnsetBlack

// Cyan wraps the content in ANSI codes to make its foreground color cyan
func Cyan(content ...interface{}) string {
	return SetCyan() + fmt.Sprint(content...) + UnsetCyan()
}

// SetCyan sets the foreground color to cyan
func SetCyan() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + cyanfg + set
}

// UnsetCyan resets the foreground color from cyan to default.
var UnsetCyan func() string = UnsetBlack

// White wraps the content in ANSI codes to make its foreground color white
func White(content ...interface{}) string {
	return SetWhite() + fmt.Sprint(content...) + UnsetWhite()
}

// SetWhite sets the foreground color to white
func SetWhite() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + whitefg + set
}

// UnsetWhite resets the foreground color from white to default.
var UnsetWhite func() string = UnsetBlack

// Color256 sets a Term256 color that is applied to the provided text
func Color256(color int, content ...interface{}) string {
	return SetColor256(color) + fmt.Sprint(content...) + UnsetColor256()
}

// SetColor256 writes the following text on the specified term256 color
func SetColor256(color int) string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + fmt.Sprintf(fg256, color) + set
}

// UnsetColor256 resets the color
var UnsetColor256 func() string = UnsetBlack

// ColorTrue sets a RGB-Color that is set as the background color, writes the
func ColorTrue(r, g, b int, content ...interface{}) string {
	return SetColorTrue(r, g, b) + fmt.Sprint(content...) + UnsetColorTrue()
}

// SetColorTrue sets a RGB-Color for the font
func SetColorTrue(r, g, b int) string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + fmt.Sprintf(fgtrue, r, g, b) + set
}

// UnsetColorTrue removes the RGB-color
var UnsetColorTrue func() string = UnsetBlack
