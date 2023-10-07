package main

import "fmt"

// MoveCursor moves the cursor to position (x;y)
func MoveCursor(x, y int) string {
	return ansiescape + fmt.Sprintf(position, x, y)
}

// Up moves the cursor up one row
func Up() string {
	return UpX(1)
}

// UpX moves the cursor up x rows
func UpX(x int) string {
	return ansiescape + fmt.Sprintf(up, x)
}

// Down moves the cursor down one row
func Down() string {
	return DownX(1)
}

// DownX moves the cursor down x rows
func DownX(x int) string {
	return ansiescape + fmt.Sprintf(down, x)
}

// Right moves the cursor right one row. As far as I am aware the correct term
// is "forward". Please be wary of potential implications in RTL context.
func Right() string {
	return RightX(1)
}

// RightX moves the cursor right x rows. As far as I am aware the correct term
// is "forward". Please be wary of potential implications in RTL context.
func RightX(x int) string {
	return ansiescape + fmt.Sprintf(forward, x)
}

// Left moves the cursor Left one row. As far as I am aware the correct term
// is "backward". Please be wary of potential implications in RTL context.
func Left() string {
	return LeftX(1)
}

// LeftX moves the cursor Left x rows. As far as I am aware the correct term
// is "backward". Please be wary of potential implications in RTL context.
func LeftX(x int) string {
	return ansiescape + fmt.Sprintf(backward, x)
}

// SavePos stores the current cursor position god knows where to allow
// restoring it.
func SavePos() string {
	return ansiescape + save
}

// RestorePos restores the cursor position from the void. (see SavePos for
// more)
func RestorePos() string {
	return ansiescape + restore
}
