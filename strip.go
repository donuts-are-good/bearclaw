package main

import (
	"bytes"
	"fmt"
)

// StripString removes all ANSI-Escape sequences from the given string and
// returns the cleaned version
//
// Links are changed so that only the URL remains; Notifications are removed
// completely.
func StripString(str string) string {
	bts := []byte(str)
	bts = stripStandard(bts)
	bts = stripLink(bts)
	bts = stripNotify(bts)

	return string(bts)
}

func stripNotify(bts []byte) []byte {
	for {
		// find the start of the notification
		index := bytes.Index(bts, []byte{0x1b, ']', '7', '7', '7', ';'})
		if index == -1 {
			break
		}
		// find the end of the notification
		removeUntil := bytes.Index(bts[index+1:], []byte{0x07})
		if removeUntil == -1 {
			break
		}

		// remove everything between start and end of the sequence
		bts = append(bts[:index], bts[index+removeUntil+2:]...)
	}

	return bts
}

func stripLink(bts []byte) []byte {
	for {
		// find the start of the url
		index := bytes.Index(bts, []byte{0x1b, '\\'})
		if index == -1 {
			break
		}
		removeUntil := bytes.Index(bts[index+1:], []byte{0x1b})
		if removeUntil == -1 {
			break
		}

		bts = append(bts[:index], bts[index+removeUntil+1:]...)
	}
	bts = bytes.ReplaceAll(bts, []byte{0x1b, ']', '8', ';', ';', 0x1b, '\\'}, []byte{})
	bts = bytes.ReplaceAll(bts, []byte{0x1b, ']', '8', ';', ';'}, []byte{})
	return bts
}

func stripStandard(bts []byte) []byte {
	matched := true
	for matched { // stop if the last run did not match anything
		index := bytes.Index(bts, []byte{0x1b, '['})
		if index == -1 {
			break
		}
		removeUntil := index + 2
		for i := removeUntil; i < len(bts); i++ {
			if (bts[i] >= 0x30 && bts[i] <= 0x39) || bts[i] == 0x3b {
				removeUntil++
				continue
			}
			if (bts[i] >= 0x41 && bts[i] <= 0x5a) || (bts[i] >= 0x61 && bts[i] <= 0x7a) {
				removeUntil++
				break
			}
			// we shouldn't be here in valid codes. anyway. stop
			// what we're doing before be break something
			fmt.Print("oppsie! ")
			fmt.Println(bts[index:removeUntil])
			matched = false
			break
		}

		bts = append(bts[:index], bts[removeUntil:]...)
	}
	return bts
}
