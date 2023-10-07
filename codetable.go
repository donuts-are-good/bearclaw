package main

const (
	// text-styles
	reset            = "0"
	bold             = "1"
	faint            = "2"
	boldOff          = "22" // also unsets faint
	italic           = "3"
	italicOff        = "23"
	underscore       = "4"
	doubleUnderscore = "21"
	underscoreOff    = "24" // removes all underscores
	blink            = "5"
	blinkOff         = "25"
	fastblink        = "6"
	fastblinkOff     = "26"
	reverseVideo     = "7" // exchanges fg and bg
	reverseVideoOff  = "27"
	conceal          = "8"
	concealOff       = "28"
	strikethrough    = "9"
	strikethroughOff = "29"
	fraktur          = "20"
	frakturOff       = "10"
	framed           = "51"
	encircled        = "52"
	framedOff        = "54"
	overlined        = "53"
	overlinedOff     = "55"

	// foregound colors
	blackfg   = "30"
	redfg     = "31"
	greenfg   = "32"
	yellowfg  = "33"
	bluefg    = "34"
	magentafg = "35"
	cyanfg    = "36"
	whitefg   = "37"
	fg256     = "38;5;%v"
	fgtrue    = "38;2;%v;%v;%v"
	resetfg   = "39"

	// background colors
	blackbg   = "40"
	redbg     = "41"
	greenbg   = "42"
	yellowbg  = "43"
	bluebg    = "44"
	magentabg = "45"
	cyanbg    = "46"
	whitebg   = "47"
	bg256     = "48;5;%v"
	bgtrue    = "48;2;%v;%v;%v"
	resetbg   = "49"

	// cursor
	position     = "%v;%vH" // line, column
	up           = "%vA"    // move cursor up x lines
	down         = "%vB"    // move cursor down x lines
	forward      = "%vC"    // move forward (usually right) x columns
	backward     = "%vD"    // move cursor backward (usually left) x columns
	save         = "s"
	restore      = "u"
	clearScreen  = "%vJ" // 0 = end of screen; 1 = beginning of screen; 2 = entire screen; 3 = clear screen and buffer
	clearLine    = "%vK" // 0 = end of line; 1 = beginning of line; 2 = entire line
	scrollUp     = "%vS" // scroll up x lines; new lines added at bottom
	scrollDown   = "%vT" // scroll down x lines; new lines added at top
	getCursorPos = "6n"

	// special
	ansiescape = "\x1b["
	set        = "m"
	escapeTC   = "\x1be["

	// hyperlink
	hyperlink = "\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\"

	// notification
	notification = "\x1b]777;notify;%s;%s\a"
)
