package main

import (
	"os"
	"syscall"
)

// EnableANSI allows windows terminals to also enjoy the beauty of formatted
// output. This is not in the init function to prevent messing with the users
// environment. On other OSes this function does absolutely nothing but does
// still exist to allow easy cross-platform development. If you are on any OS
// other than Windows and your Terminal does not support ANSI Sequences I'd
// reccomend you get a normal terminal emulator.
func EnableANSI() {
	handle := syscall.Handle(os.Stdout.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
	setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004)
}
