package main

import "fmt"

// ClearRight clears all characters right of the cursor from the display
func ClearRight() string {
	return ansiescape + fmt.Sprintf(clearLine, 0)
}

// ClearLeft clears all characters left of the cursor from the display
func ClearLeft() string {
	return ansiescape + fmt.Sprintf(clearLine, 1)
}

// ClearLine clears all characters on the line
func ClearLine() string {
	return ansiescape + fmt.Sprintf(clearLine, 2)
}

// ClearScreen clears the entire screen and moves the cursor to (0;0)
func ClearScreen() string {
	return ansiescape + fmt.Sprintf(clearScreen, 2)
}

// ClearBegin clears the entire screen to (0;0)
func ClearBegin() string {
	return ansiescape + fmt.Sprintf(clearScreen, 1)
}

// ClearEnd clears to the end of the screen
func ClearEnd() string {
	return ansiescape + fmt.Sprintf(clearScreen, 0)
}

// ClearScreenBuffer clears the entire screen and the screenbuffer
func ClearScreenBuffer() string {
	return ansiescape + fmt.Sprintf(clearScreen, 3)
}
