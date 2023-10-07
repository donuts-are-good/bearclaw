package main

import "fmt"

// BlackBG wraps the content in ANSI codes to make its background color black
func BlackBG(content ...interface{}) string {
	return SetBlackBG() + fmt.Sprint(content...) + UnsetBlackBG()
}

// SetBlackBG sets the background color to black
func SetBlackBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + blackbg + set
}

// UnsetBlackBG resets the background color from black to default.
func UnsetBlackBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + resetbg + set
}

// RedBG wraps the content in ANSI codes to make its background color red
func RedBG(content ...interface{}) string {
	return SetRedBG() + fmt.Sprint(content...) + UnsetRedBG()
}

// SetRedBG sets the background color to red
func SetRedBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + redbg + set
}

// UnsetRedBG resets the background color from red to default.
var UnsetRedBG func() string = UnsetBlackBG

// GreenBG wraps the content in ANSI codes to make its background color green
func GreenBG(content ...interface{}) string {
	return SetGreenBG() + fmt.Sprint(content...) + UnsetGreenBG()
}

// SetGreenBG sets the background color to green
func SetGreenBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + greenbg + set
}

// UnsetGreenBG resets the background color from green to default.
var UnsetGreenBG func() string = UnsetBlackBG

// YellowBG wraps the content in ANSI codes to make its background color yellow
func YellowBG(content ...interface{}) string {
	return SetYellowBG() + fmt.Sprint(content...) + UnsetYellowBG()
}

// SetYellowBG sets the background color to yellow
func SetYellowBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + yellowbg + set
}

// UnsetYellowBG resets the background color from yellow to default.
var UnsetYellowBG func() string = UnsetBlackBG

// BlueBG wraps the content in ANSI codes to make its background color blue
func BlueBG(content ...interface{}) string {
	return SetBlueBG() + fmt.Sprint(content...) + UnsetBlueBG()
}

// SetBlueBG sets the background color to blue
func SetBlueBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + bluebg + set
}

// UnsetBlueBG resets the background color from blue to default.
var UnsetBlueBG func() string = UnsetBlackBG

// MagentaBG wraps the content in ANSI codes to make its background color magenta
func MagentaBG(content ...interface{}) string {
	return SetMagentaBG() + fmt.Sprint(content...) + UnsetMagentaBG()
}

// SetMagentaBG sets the background color to magenta
func SetMagentaBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + magentabg + set
}

// UnsetMagentaBG resets the background color from magenta to default.
var UnsetMagentaBG func() string = UnsetBlackBG

// CyanBG wraps the content in ANSI codes to make its background color cyan
func CyanBG(content ...interface{}) string {
	return SetCyanBG() + fmt.Sprint(content...) + UnsetCyanBG()
}

// SetCyanBG sets the background color to cyan
func SetCyanBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + cyanbg + set
}

// UnsetCyanBG resets the background color from cyan to default.
var UnsetCyanBG func() string = UnsetBlackBG

// WhiteBG wraps the content in ANSI codes to make its background color white
func WhiteBG(content ...interface{}) string {
	return SetWhiteBG() + fmt.Sprint(content...) + UnsetWhiteBG()
}

// SetWhiteBG sets the background color to white
func SetWhiteBG() string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + whitebg + set
}

// UnsetWhiteBG resets the background color from white to default.
var UnsetWhiteBG func() string = UnsetBlackBG

// Color256BG sets a Term256 color that is applied to the provided text as the
// background color
func Color256BG(color int, content ...interface{}) string {
	return SetColor256BG(color) + fmt.Sprint(content...) + UnsetColor256BG()
}

// SetColor256BG writes the following text on the specified background color
func SetColor256BG(color int) string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + fmt.Sprintf(bg256, color) + set
}

// UnsetColor256BG resets the background color
var UnsetColor256BG func() string = UnsetBlackBG

// ColorTrueBG sets a RGB-Color that is set as the background color, writes the
// text and clears the background color
func ColorTrueBG(r, g, b int, content ...interface{}) string {
	return SetColorTrueBG(r, g, b) + fmt.Sprint(content...) + UnsetColorTrueBG()
}

// SetColorTrueBG sets a RGB-Color for the background
func SetColorTrueBG(r, g, b int) string {
	if nocolorIsSet {
		return ""
	}
	return ansiescape + fmt.Sprintf(bgtrue, r, g, b) + set
}

// UnsetColorTrueBG removes the RGB-background
var UnsetColorTrueBG func() string = UnsetBlackBG
