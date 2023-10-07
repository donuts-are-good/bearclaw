package main

import "os"

var nocolorIsSet bool

func init() {
	_, nocolorIsSet = os.LookupEnv("NO_COLOR")
}

// OverrideNoColor allows re-enabling color codes if the user explicitly enabled
// it by means of a configuration entry or commandline flag.
func OverrideNoColor() {
	nocolorIsSet = false
}
